package presenters

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type ChatRequest struct {
	Input string `json:"input" validate:"required"`
}

func ChatErrorResponse(title string, err error) *fiber.Map {
	return &fiber.Map{
		"title": title,
		"error": err.Error(),
	}
}
func ChatResponse(chat *entities.Chat) *fiber.Map {
	return &fiber.Map{
		"id":          chat.ID,
		"workflow_id": chat.WorkflowID,
		"created_at":  chat.CreatedAt,
		"updated_at":  chat.UpdatedAt,
	}
}
