package routes

import (
	"github.com/Ratchaphon1412/assistant-llm/api/handlers"
	"github.com/Ratchaphon1412/assistant-llm/api/middlewares"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/account"

	"github.com/gofiber/fiber/v2"
)

func AccountRouter(app fiber.Router, auth_middleware fiber.Handler, db database.Dbinstance, cfg *configs.Config) {
	accountService := account.NewService(account.NewRepository(db.Db))
	app.Get("/auth/google", middlewares.LoggerRequest(cfg), handlers.GoogleSignIn(accountService, cfg))
	app.Get("/auth/google/callback", middlewares.LoggerRequest(cfg), handlers.GoogleCallback(accountService, cfg))
	app.Get("/account", middlewares.LoggerRequest(cfg), auth_middleware, middlewares.ExtractToken, handlers.GetAccount(accountService))
	app.Get("/account/logout", middlewares.LoggerRequest(cfg), auth_middleware, middlewares.ExtractToken, handlers.Logout(accountService, cfg))
}
