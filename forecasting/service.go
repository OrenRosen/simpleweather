package forecasting

import (
	"fmt"

	"github.com/OrenRosen/simpleweather/weather"
)

type WeatherProvider interface {
	GetWeatherByCity(city string) (weather.Weather, error)
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
	w, err := s.weatherProvider.GetWeatherByCity(city)
	if err != nil {
		return "", fmt.Errorf("WhatToWear: %w", err)
	}

	if w.Temp < 21 {
		return "long sleeves", nil
	}

	return "short sleeves", nil
}
