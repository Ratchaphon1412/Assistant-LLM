package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ratchaphon1412/assistant-llm/configs"
)

type Service interface {
	CurrentWeather(lat string, lon string) (*CurrentWeatherData, error)
	ForCastWeather(lat string, lon string) (*ForecastWeatherData, error)
}

type CurrentWeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  *int    `json:"sea_level"`
		GrndLevel *int    `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64  `json:"speed"`
		Deg   int      `json:"deg"`
		Gust  *float64 `json:"gust,omitempty"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type ForecastWeatherData struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Rain       struct {
			OneH float64 `json:"1h"`
		} `json:"rain,omitempty"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int64  `json:"sunrise"`
		Sunset     int64  `json:"sunset"`
	} `json:"city"`
}

type service struct {
	cfg *configs.Config
}

func NewService(cfg *configs.Config) Service {
	return &service{
		cfg: cfg,
	}
}

// CurrentWeather implements Service.
func (s *service) CurrentWeather(lat string, lon string) (*CurrentWeatherData, error) {

	apiURL := fmt.Sprintf("%s/data/2.5/weather?lat=%s&lon=%s&units=metric&lang=th&appid=%s",
		s.cfg.OPEN_WEATHER_MAP_URL, lat, lon, s.cfg.OPEN_WEATHER_MAP_API_KEY,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data CurrentWeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// You can use 'data' as needed here

	return &data, nil
}

func (s *service) ForCastWeather(lat string, lon string) (*ForecastWeatherData, error) {
	apiURL := fmt.Sprintf("%s/data/2.5/forecast?lat=%s&lon=%s&units=metric&lang=th&appid=%s",
		s.cfg.OPEN_WEATHER_MAP_URL, lat, lon, s.cfg.OPEN_WEATHER_MAP_API_KEY,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data ForecastWeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}
