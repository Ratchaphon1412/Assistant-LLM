package handlers

import (
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/auth"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"
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

		return c.Status(200).JSON(fiber.Map{"email": userinfo.Email, "profile": userinfo.Picture, "login": true})
	}

}
