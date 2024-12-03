# **Deploy-CloudRun-GO**

Este projeto é uma aplicação desenvolvida em Go para buscar informações climáticas com base em um CEP fornecido pelo usuário. Ele utiliza serviços externos, como **ViaCEP** e **WeatherAPI**, e é configurado para ser implantado no **Google Cloud Run**.

---

## **Link Acesso Cloud Run**

https://deploy-com-clound-go-453723015584.us-central1.run.app/weather?cep=20081000

## **Funcionalidades**

- Recebe um **CEP** válido (8 dígitos) e identifica a cidade correspondente.
- Faz chamadas à API **ViaCEP** para buscar a localização e à **WeatherAPI** para obter informações climáticas.
- Retorna as temperaturas nas escalas **Celsius (°C)**, **Fahrenheit (°F)** e **Kelvin (K)**.
- Garante respostas adequadas para casos de sucesso e falha.

---

## **Requisitos**

- **Go 1.23+**
- **Docker** e **Docker Compose**
- **Google Cloud SDK** (gcloud)
- **Skaffold** (para fluxos de build e deploy no Cloud Run)

---

## **Configuração do Projeto**

1. **Clone o Repositório**
   git clone https://github.com/seu-repositorio/Deploy-CloudRun-GO.git
   cd Deploy-CloudRun-GO
