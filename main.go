package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/OrenRosen/simpleweather/forecasting"
	"github.com/OrenRosen/simpleweather/openweather"
)

const apiKey = "put_your_api_key_here"

func main() {
	cityP := flag.String("city", "London", "City to be queried")
	flag.Parse()
	city := *cityP

	openweatherProvider := openweather.NewProvider(apiKey)
	weatherService := forecasting.NewService(openweatherProvider)

	outfit, err := weatherService.WhatToWear(city)
	if err != nil {
		log.Fatalf("couldn't get what to wear in %s: %v", city, err)
	}

	fmt.Printf("you should wear in %s: %s\n", city, outfit)
}
