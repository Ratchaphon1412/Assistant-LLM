package handlers

import (
	"errors"

	"time"

	"github.com/Ratchaphon1412/assistant-llm/api/presenters"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type GoogleCallBackRequest struct {
	Code string `json:"code"`
}

func GoogleSignIn(service account.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url, err := service.GoogleSignIn(cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get google sign in url", err))
		}
		return c.Redirect(url)
	}
}

func GoogleCallback(service account.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AccountErrorResponse("required", errors.New("Code is required")))
		}
		userinfo, err := service.GoogleCallback(c.Context(), code, cfg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get user info", err))
		}

		account, err := service.GetAccountByEmail(userinfo.Email)
		if account == nil {
			account = &entities.Account{
				Email:   userinfo.Email,
				Profile: userinfo.Picture,
			}
			account, err = service.CreateAccount(account)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to create account", err))
			}
			t, err := service.SignIn(account, *cfg)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to sign in", err))
			}

			// Set JWT as HTTP Only Cookie
			c.Cookie(&fiber.Cookie{
				Name:     cfg.JWT_COOKIE_NAME,
				Value:    t,
				Path:     "/",                            // ให้ cookie ใช้ได้ทั้งเว็บ
				HTTPOnly: cfg.JWT_HTTP_ONLY,              // ป้องกัน JS access
				Secure:   cfg.JWT_SECURE,                 // ควรใช้ production (HTTPS เท่านั้น)
				SameSite: "Lax",                          // ปรับเป็น "None" ถ้ามี cross-site
				Expires:  time.Now().Add(time.Hour * 72), // ตั้งอายุ cookie ตามต้องการ
			})
			return c.Redirect(cfg.CLIENT_URL+"/chat", fiber.StatusSeeOther)

		} else {
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get account", err))
			}

			t, err := service.SignIn(account, *cfg)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to sign in", err))

			}
			c.Cookie(&fiber.Cookie{
				Name:     cfg.JWT_COOKIE_NAME,
				Value:    t,
				Path:     "/",                            // ให้ cookie ใช้ได้ทั้งเว็บ
				HTTPOnly: cfg.JWT_HTTP_ONLY,              // ป้องกัน JS access
				Secure:   cfg.JWT_SECURE,                 // ควรใช้ production (HTTPS เท่านั้น)
				SameSite: "Lax",                          // ปรับเป็น "None" ถ้ามี cross-site
				Expires:  time.Now().Add(time.Hour * 72), // ตั้งอายุ cookie ตามต้องการ
			})

			return c.Redirect(cfg.CLIENT_URL+"/chat", fiber.StatusSeeOther)

		}

	}

}

func GetAccount(service account.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		account, err := service.GetAccountByEmail(email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get account", err))
		}
		return c.Status(200).JSON(presenters.AccountResponse(account))
	}
}

func Logout(service account.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)
		account, err := service.GetAccountByEmail(email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get account", err))
		}
		if account == nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.AccountErrorResponse("Account not found", errors.New("Account not found")))
		}

		// Clear JWT Cookie
		if c.Cookies(cfg.JWT_COOKIE_NAME, "") == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AccountErrorResponse("Cookie not found", errors.New("Cookie not found")))
		}

		c.Cookie(&fiber.Cookie{
			Name:    cfg.JWT_COOKIE_NAME,
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		})
		return c.Status(fiber.StatusOK).JSON(presenters.AccountResponse(account))

	}
}
