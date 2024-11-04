FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o weather-api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/weather-api /app/weather-api

EXPOSE 8080

CMD ["/app/weather-api"]

