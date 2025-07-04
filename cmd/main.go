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

	if err := env.Parse(&cfg); err != nil {
		panic(err) //TODO: write to log
	}
	configs.AppSettings(app, &cfg)
	// Initialize Database
	database.Connect(&cfg)

	// Initialize Redis
	database.ConnectRedis(&cfg)

	api := app.Group("/api")
	websocket := app.Group("/ws")

	// middleware
	auth_middleware := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(cfg.JWT_SECRET)},
		TokenLookup: "cookie:" + cfg.JWT_COOKIE_NAME,
	})

	//api v1
	apiV1 := api.Group("/v1")
	websocketV1 := websocket.Group("/v1")

	routes.AccountRouter(apiV1, auth_middleware, database.DB, &cfg)
	routes.ChatRouter(websocketV1, auth_middleware, database.DB, &cfg)
	routes.WeatherRouter(apiV1, auth_middleware, database.DB, &cfg)
	app.Listen("0.0.0.0:" + cfg.ServerPort)
}
