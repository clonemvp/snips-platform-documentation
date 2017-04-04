package openweathermap

import (
	"net/http"
	"fmt"
	"snips-platform-lambda-samples/go/weather/conf"
	"io/ioutil"
	"encoding/json"
	"log"
)

func GetWeather(wc conf.WeatherConf, city string) string {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%v&APPID=%v&units=metric", city, wc.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error when retrieving weather: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error when reading weather: %v", err)
	}

	var owm OwmResponse
	if err := json.Unmarshal(body, &owm); err != nil {
		log.Fatalf("json.Unmarshal, err: %v", err)
	}
	log.Printf("owm: %+v", owm)
	temp := int(owm.Main.Temp)

	return fmt.Sprintf("The temperature is %d degrees in %v", temp, city)
}

type OwmResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}