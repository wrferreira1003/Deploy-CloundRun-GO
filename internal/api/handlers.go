package api

import (
	"encoding/json"
	"net/http"

	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/models"
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/usecase"
)

type WeatherHandler struct {
	weatherUsecase usecase.WeatherUsecaseIn
}

func NewWeatherHandler(weatherUsecase usecase.WeatherUsecaseIn) *WeatherHandler {
	return &WeatherHandler{
		weatherUsecase: weatherUsecase,
	}
}

func (h *WeatherHandler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	// Buscar o clima
	weather, err := h.weatherUsecase.GetWeatherByCep(cep)
	if err != nil {
		h.handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

func (h *WeatherHandler) handleError(w http.ResponseWriter, err error) {
	var statusCode int
	var message string

	switch err {
	case models.ErrInvalidZipCode:
		statusCode = http.StatusUnprocessableEntity
		message = `{"message": "invalid zipcode"}`
	case models.ErrZipCodeNotFound:
		statusCode = http.StatusNotFound
		message = `{"message": "can not find zipcode"}`
	case models.ErrWeatherNotFound:
		statusCode = http.StatusNotFound
		message = `{"message": "can not find weather"}`
	default:
		statusCode = http.StatusInternalServerError
		message = `{"message": "internal server error"}`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
