package handlers

import (
	"net/http"

	"github.com/belmadge/temperature-cep/services"

	"github.com/gin-gonic/gin"
)

func WeatherHandler(c *gin.Context) {
	cep := c.Param("cep")
	if len(cep) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	location, err := services.GetLocationByCEP(cep)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	tempC, err := services.GetTemperature(location.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve temperature"})
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	c.JSON(http.StatusOK, gin.H{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	})
}
