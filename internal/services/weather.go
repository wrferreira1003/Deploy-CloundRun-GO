package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherServiceInterface interface {
	GetTemperature(city string) (float64, float64, error)
}

type WeatherAPIService struct {
	APIKey  string
	BaseURL string // Adicionando URL base configurável
}

func NewWeatherAPIService(apiKey string, baseURL string) *WeatherAPIService {
	return &WeatherAPIService{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (s *WeatherAPIService) GetTemperature(city string) (float64, float64, error) {
	// Codificar o nome da cidade
	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("%s?key=%s&q=%s", s.BaseURL, s.APIKey, encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("weatherapi returned status: %d", resp.StatusCode)
	}

	var data struct {
		Current struct {
			TempC float64 `json:"temp_c"`
			TempF float64 `json:"temp_f"`
		} `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, fmt.Errorf("failed to parse weather data: %w", err)
	}

	return data.Current.TempC, data.Current.TempF, nil
}
