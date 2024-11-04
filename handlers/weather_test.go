package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWeatherHandlerInvalidCEP(t *testing.T) {
	router := gin.Default()
	router.GET("/weather/:cep", WeatherHandler)

	req, _ := http.NewRequest("GET", "/weather/123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, got %d", resp.Code)
	}
}
