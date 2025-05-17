package routes

import (
	"github.com/Ratchaphon1412/assistant-llm/api/handlers"
	"github.com/Ratchaphon1412/assistant-llm/api/middlewares"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/chat"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app fiber.Router, auth_middleware fiber.Handler, db database.Dbinstance, cfg *configs.Config) {
	chatService := chat.NewService(chat.NewRepository(db.Db))
	app.Get("/chat", auth_middleware, middlewares.ExtractToken, middlewares.UpgradeRequest, handlers.CreateChat(chatService, cfg))

}
