package handlers

import (
	"errors"

	"github.com/Ratchaphon1412/assistant-llm/api/presenters"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/weather"
	"github.com/gofiber/fiber/v2"
)

func GetCurrentWeather(service weather.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		lat := c.Query("lat")
		lon := c.Query("lon")
		if lat == "" || lon == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.WeatherErrorResponse("Latitude and Longitude are required", errors.New("Latitude or Longitude are required")))

		}
		weatherData, err := service.CurrentWeather(lat, lon)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.WeatherErrorResponse("Failed to get current weather", err))
		}

		return c.Status(fiber.StatusOK).JSON(presenters.CurrentWeatherResponse(weatherData))
	}
}

func GetWeatherForecast(service weather.Service, cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		lat := c.Query("lat")
		lon := c.Query("lon")
		if lat == "" || lon == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.WeatherErrorResponse("Latitude and Longitude are required", errors.New("Latitude or Longitude are required")))
		}
		forecastData, err := service.ForCastWeather(lat, lon)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.WeatherErrorResponse("Failed to get weather forecast", err))
		}

		return c.Status(fiber.StatusOK).JSON(presenters.ForecastWeatherResponse(forecastData))
	}
}
