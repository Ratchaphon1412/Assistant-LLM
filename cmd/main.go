package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Ratchaphon1412/assistant-llm/api/routes"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"

	"github.com/caarlos0/env/v11"
	jwtware "github.com/gofiber/contrib/jwt"
	_ "github.com/joho/godotenv/autoload"
)

var cfg configs.Config

func main() {
	app := fiber.New()
	configs.AppSettings(app)
	if err := env.Parse(&cfg); err != nil {
		panic(err) //TODO: write to log
	}

	database.Connect(&cfg)
	api := app.Group("/api")

	// middleware
	auth_middleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWT_SECRET)},
	})

	//api v1
	v1 := api.Group("/v1")

	routes.AccountRouter(v1, auth_middleware, database.DB, &cfg)
	app.Listen(":" + cfg.ServerPort)
}
