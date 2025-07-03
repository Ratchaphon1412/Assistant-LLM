package routes

import (
	"github.com/Ratchaphon1412/assistant-llm/api/handlers"
	"github.com/Ratchaphon1412/assistant-llm/api/middlewares"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/weather"
	"github.com/gofiber/fiber/v2"
)

func WeatherRouter(app fiber.Router, auth_middleware fiber.Handler, db database.Dbinstance, cfg *configs.Config) {
	wetherService := weather.NewService(cfg)
	app.Get("/weather", middlewares.LoggerRequest(cfg), auth_middleware, middlewares.ExtractToken, handlers.GetCurrentWeather(wetherService, cfg))
	app.Get("/weather/forecast", middlewares.LoggerRequest(cfg), auth_middleware, middlewares.ExtractToken, handlers.GetWeatherForecast(wetherService, cfg))

}
