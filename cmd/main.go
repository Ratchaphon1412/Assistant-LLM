package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Ratchaphon1412/assistant-llm/api/routes"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var cfg configs.Config

func main() {
	app := fiber.New()
	if err := env.Parse(&cfg); err != nil {
		panic(err) //TODO: write to log
	}

	database.Connect(&cfg)
	api := app.Group("/api")
	routes.AccountRouter(api, database.DB, &cfg)
	app.Listen(":" + cfg.ServerPort)
}
