package presenters

import (
	"github.com/Ratchaphon1412/assistant-llm/pkg/weather"
	"github.com/gofiber/fiber/v2"
)

func WeatherErrorResponse(title string, err error) *fiber.Map {
	return &fiber.Map{
		"title": title,
		"error": err.Error(),
	}
}
func CurrentWeatherResponse(weather *weather.CurrentWeatherData) *fiber.Map {
	return &fiber.Map{
		"temp":                weather.Main.Temp,
		"weather_name":        weather.Weather[0].Main,
		"weather_description": weather.Weather[0].Description,
		"icon":                weather.Weather[0].Icon,
		"location_name":       weather.Name,
	}
}

func ForecastWeatherResponse(weather *weather.ForecastWeatherData) []map[string]interface{} {
	responses := make([]map[string]interface{}, 0, len(weather.List))
	for _, item := range weather.List {
		if len(item.Weather) == 0 {
			continue
		}
		responses = append(responses, map[string]interface{}{
			"temp":                item.Main.Temp,
			"weather_name":        item.Weather[0].Main,
			"weather_description": item.Weather[0].Description,
			"icon":                item.Weather[0].Icon,
			"date_txt":            item.DtTxt,
		})
	}
	return responses

}
