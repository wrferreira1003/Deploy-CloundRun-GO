package usecase

import (
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/models"
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/services"
)

type WeatherUsecaseIn interface {
	GetWeatherByCep(cep string) (models.TemperatureResponse, error)
}

type weatherUsecase struct {
	weatherService  services.WeatherServiceInterface
	locationService services.LocationServiceInterface
}

func NewWeatherUsecase(
	weatherService services.WeatherServiceInterface,
	locationService services.LocationServiceInterface,
) *weatherUsecase {
	return &weatherUsecase{
		weatherService:  weatherService,
		locationService: locationService,
	}
}

func (w *weatherUsecase) GetWeatherByCep(cep string) (models.TemperatureResponse, error) {
	if len(cep) != 8 {
		return models.TemperatureResponse{}, models.ErrInvalidZipCode
	}

	//Buscar o cep
	location, err := w.locationService.GetLocationByCep(cep)
	if err != nil {
		return models.TemperatureResponse{}, models.ErrZipCodeNotFound
	}

	//Buscar o clima
	tempC, tempF, err := w.weatherService.GetTemperature(location)
	if err != nil {
		return models.TemperatureResponse{}, models.ErrWeatherNotFound
	}

	tempK := tempC + 273.15

	return models.TemperatureResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
