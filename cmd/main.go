package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wrferreira1003/Deploy-Cloud-GO/configs"
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/api"
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/services"
	"github.com/wrferreira1003/Deploy-Cloud-GO/internal/usecase"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	fmt.Println("Config loaded successfully", cfg)

	weatherService := services.NewWeatherAPIService(cfg.WeatherApiKey, cfg.WeatherBaseURL)
	locationService := services.NewViaCepService(cfg)

	weatherUsecase := usecase.NewWeatherUsecase(weatherService, locationService)
	weatherHandler := api.NewWeatherHandler(weatherUsecase)

	http.HandleFunc("/weather", weatherHandler.GetWeatherHandler)

	http.ListenAndServe(":8080", nil)
}
