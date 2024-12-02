package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherAPIService_GetTemperature(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Get("key") != "mock-api-key" {
			http.Error(w, "missing or incorrect API key", http.StatusUnauthorized)
			return
		}
		if query.Get("q") != "São Paulo" {
			http.Error(w, "city not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"current": {"temp_c": 28.5, "temp_f": 83.3}}`)
	}))
	defer mockServer.Close()

	// Configurar o serviço para usar o mockServer
	mockService := &WeatherAPIService{
		APIKey:  "mock-api-key",
		BaseURL: mockServer.URL, // Redirecionar chamadas para o mockServer
	}

	t.Run("Sucesso", func(t *testing.T) {
		tempC, tempF, err := mockService.GetTemperature("São Paulo")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if tempC != 28.5 || tempF != 83.3 {
			t.Errorf("expected temps: 28.5°C, 83.3°F; got: %v°C, %v°F", tempC, tempF)
		}
	})

	t.Run("Erro - Cidade não encontrada", func(t *testing.T) {
		_, _, err := mockService.GetTemperature("Cidade Inexistente")
		if err == nil || err.Error() != "weatherapi returned status: 404" {
			t.Errorf("expected error 'weatherapi returned status: 404', got %v", err)
		}
	})

	t.Run("Erro - Chave de API inválida", func(t *testing.T) {
		mockService.APIKey = "invalid-api-key"

		_, _, err := mockService.GetTemperature("São Paulo")
		if err == nil || err.Error() != "weatherapi returned status: 401" {
			t.Errorf("expected error 'weatherapi returned status: 401', got %v", err)
		}
	})
}
