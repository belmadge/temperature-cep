package main

import (
	"log"
	"os"

	"github.com/belmadge/temperature-cep/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/weather/:cep", handlers.WeatherHandler)

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}
