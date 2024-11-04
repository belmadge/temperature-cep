package main

import (
	"log"

	"github.com/belmadge/temperature-cep/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/weather/:cep", handlers.WeatherHandler)
	router.Run(":8080")
}
