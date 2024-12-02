FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=builder /app/cloudrun .
ENTRYPOINT ["./cloudrun"]
