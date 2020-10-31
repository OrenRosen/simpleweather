package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OrenRosen/simpleweather/temprature"
)

const (
	endpoint                = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByCity = "/weather?q=%s&appid=%s&units=metric"
)

type provider struct {
	apiKey string
}

func NewProvider(apiKey string) *provider {
	return &provider{
		apiKey: apiKey,
	}
}

func (p *provider) GetWeatherByCity(city string) (temprature.Weather, error) {
	path := fmt.Sprintf(pathFormatWeatherByCity, city, p.apiKey)
	u := endpoint + path

	res, err := http.Get(u)
	if err != nil {
		return temprature.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return temprature.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed reading body: %s", err)
	}

	var weatherRes weatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherRes); err != nil {
		return temprature.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return temprature.Weather{}, fmt.Errorf("openweather.GetWeatherByCity got error from OpenWeather: %s", weatherRes.Message)
	}

	return weatherRes.ToWeather(), nil
}

type weatherResponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	}
}

func (r weatherResponse) ToWeather() temprature.Weather {
	return temprature.Weather{
		Temp:     r.Main.Temp,
		Pressure: r.Main.Pressure,
		MinTemp:  r.Main.TempMin,
		MaxTemp:  r.Main.TempMax,
	}
}
