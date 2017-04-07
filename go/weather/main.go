package main

import (
	"os"
	"snips-platform-lambda-samples/go/common/mqtt"
	"snips-platform-lambda-samples/go/weather/conf"
	"encoding/json"
	"snips-platform-lambda-samples/go/weather/openweathermap"
	"strings"
	"snips-platform-lambda-samples/go/common/models"
	"log"
)

func main() {
	conf.Run(os.Args, func(wc conf.WeatherConf) {
		client := mqtt.NewClient(wc.PlatformConf)
		client.Connect()
		defer client.Disconnect()

		cityChannel := make(chan string)

		client.Subscribe("hermes/Intent/GetWeather", func(topic string, payload []byte) {
			intent := models.ParseIntent(payload)
			probability := intent.Intent.Probability
			log.Printf("received get weather [%v] with probability: %v", intent, probability)
			if probability > 0.75 && len(intent.Slots) > 0 {
				value := intent.Slots[0].Value
				cityChannel <- strings.ToLower(value)
			}
		})

		for {
			m := make(map[string]string)
			m["text"] = openweathermap.GetWeather(wc, <-cityChannel)
			message, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("main: %v", err)
			}
			client.Publish("hermes/Tts", message)
		}
	})
}
