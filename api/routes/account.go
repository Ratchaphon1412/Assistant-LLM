package routes

import (
	"github.com/Ratchaphon1412/assistant-llm/api/handlers"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"
	"github.com/gofiber/fiber/v2"
)

func AccountRouter(app fiber.Router, db database.Dbinstance, cfg *configs.Config) {
	accountService := account.NewService(account.NewRepository(db.Db))
	app.Get("/auth/google", handlers.GoogleSignIn(accountService, cfg))
	app.Get("/auth/google/callback", handlers.GoogleCallback(accountService, cfg))
}
