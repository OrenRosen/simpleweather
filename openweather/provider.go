package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OrenRosen/simpleweather/weather"
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

func (p *provider) GetWeatherByCity(city string) (weather.Weather, error) {
	// compose the url. note that it's not the best way to add query params.
	path := fmt.Sprintf(pathFormatWeatherByCity, city, p.apiKey)
	u := endpoint + path

	res, err := http.Get(u)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// read the response body and encode it into the respose struct
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed reading body: %s", err)
	}

	var weatherRes weatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherRes); err != nil {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByCity failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return weather.Weather{}, fmt.Errorf("openweather.GetWeatherByCity got error from OpenWeather: %s", weatherRes.Message)
	}

	// return the external response converted into an entity
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

func (r weatherResponse) ToWeather() weather.Weather {
	return weather.Weather{
		Temp:     r.Main.Temp,
		Pressure: r.Main.Pressure,
		MinTemp:  r.Main.TempMin,
		MaxTemp:  r.Main.TempMax,
	}
}
