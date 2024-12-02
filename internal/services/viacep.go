package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/wrferreira1003/Deploy-Cloud-GO/configs"
)

// Interface para o serviço de localização
type LocationServiceInterface interface {
	GetLocationByCep(cep string) (string, error)
}

// Implementação do serviço usando ViaCEP
type ViaCepService struct {
	cfg *configs.Config
}

func NewViaCepService(cfg *configs.Config) *ViaCepService {
	return &ViaCepService{
		cfg: cfg,
	}
}

func (v *ViaCepService) GetLocationByCep(cep string) (string, error) {
	url := fmt.Sprintf("%s/%s/json/", v.cfg.ViaCepUrl, cep)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("error fetching data from ViaCep")
	}
	defer resp.Body.Close()

	var data struct {
		Localidade string `json:"localidade"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", errors.New("error decoding response from ViaCep")
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("location not found for cep: %s", cep)
	}

	return data.Localidade, nil
}
