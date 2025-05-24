package handlers

import (
	"github.com/Ratchaphon1412/assistant-llm/api/presenters"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/fiber/v2"
	// "errors"
)

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
		userinfo, err := service.GoogleCallback(c.Context(), c.Query("code"), cfg)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AccountErrorResponse("Failed to get user info", err))
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
			return c.Status(fiber.StatusOK).JSON(presenters.SignGoogleCallBackResponse(account, t))
		} else {
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to get account", err))
			}

			t, err := service.SignIn(account, *cfg)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AccountErrorResponse("Failed to sign in", err))

			}
			return c.Status(fiber.StatusOK).JSON(presenters.SignGoogleCallBackResponse(account, t))

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
