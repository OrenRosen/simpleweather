package forecasting

import (
	"fmt"

	"github.com/OrenRosen/simpleweather/temprature"
)

type WeatherProvider interface {
	GetWeatherByCity(city string) (temprature.Weather, error)
}

type service struct {
	weatherProvider WeatherProvider
}

func NewService(p WeatherProvider) *service {
	return &service{
		weatherProvider: p,
	}
}

func (s *service) WhatToWear(city string) (string, error) {
	weather, err := s.weatherProvider.GetWeatherByCity(city)
	if err != nil {
		return "", fmt.Errorf("WhatToWear: %w", err)
	}

	if weather.Temp < 24 {
		return "long sleeves", nil
	}

	return "short sleeves", nil
}
