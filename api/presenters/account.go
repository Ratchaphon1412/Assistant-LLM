package presenters

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AccountErrorResponse(title string, err error) *fiber.Map {
	return &fiber.Map{
		"title": title,
		"error": err.Error(),
	}
}

func SignGoogleCallBackResponse(account *entities.Account, token string) *fiber.Map {
	return &fiber.Map{
		"email":   account.Email,
		"profile": account.Profile,
		"token":   token,
	}
}
func AccountResponse(account *entities.Account) *fiber.Map {
	return &fiber.Map{
		"email":   account.Email,
		"profile": account.Profile,
	}
}
