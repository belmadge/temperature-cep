FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o weather-api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/weather-api /app/weather-api
COPY .env /app/.env

EXPOSE 8080

CMD ["/app/weather-api"]
