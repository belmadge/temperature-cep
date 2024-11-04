package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Location struct {
	City string `json:"localidade"`
}

func GetLocationByCEP(cep string) (*Location, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("CEP not found")
	}

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.City == "" {
		return nil, errors.New("invalid location data")
	}

	return &location, nil
}
