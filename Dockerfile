# Etapa de construção
FROM golang:1.23.1 as builder

WORKDIR /app

# Copiar arquivos e construir o binário
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main.go

# Etapa final com Debian Slim
FROM debian:bullseye-slim
WORKDIR /app

# Instalar certificados CA
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copiar o binário e arquivos necessários
COPY --from=builder /app/cloudrun .
COPY --from=builder /app/.env .

# Comando de execução
ENTRYPOINT ["./cloudrun"]



