package handlers

import (
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/auth"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/fiber/v2"
	// "errors"
)

func GoogleSignIn(service account.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := auth.ConfigGoogle(cfg)
		url := path.AuthCodeURL("state")
		return c.Redirect(url)
	}
}

func GoogleCallback(service account.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, error := auth.ConfigGoogle(cfg).Exchange(c.Context(), c.FormValue("code"))
		if error != nil {
			panic(error)
		}
		userinfo := auth.GetUserInfo(token.AccessToken)
		if userinfo.Email == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid email"})
		}
		account, err := service.GetAccountByEmail(userinfo.Email)
		if account == nil {
			account = &entities.Account{
				Email:   userinfo.Email,
				Profile: userinfo.Picture,
			}
			account, err = service.CreateAccount(account)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to create account ", "error_message": err.Error()})
			}
			t, err := service.SignIn(account, *cfg)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to sign in", "error_message": err.Error(), "user_info": userinfo})
			}
			return c.Status(200).JSON(fiber.Map{"email": account.Email, "profile": account.Profile, "login": true, "token": t})
		} else {
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to get account", "error_message": err.Error()})
			}

			t, err := service.SignIn(account, *cfg)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to sign in", "error_message": err.Error()})

			}
			return c.Status(200).JSON(fiber.Map{"email": account.Email, "profile": account.Profile, "login": true, "token": t})

		}
	}

}
