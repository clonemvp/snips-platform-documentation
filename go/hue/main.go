package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"snips-platform-lambda-samples/go/hue/mqtt"
	"snips-platform-lambda-samples/go/hue/conf"
	"snips-platform-lambda-samples/go/hue/bulb"
)

func main() {
	conf.Run(os.Args, func(sc conf.SampleConf) {
		client := mqtt.NewClient(sc.PlatformHost, sc.PlatformPort)
		client.Connect()
		defer client.Disconnect()

		stateChannel := make(chan bulb.State)

		client.Subscribe("hermes/Intent/set_light_color", func(topic string, payload []byte) {
			value := parse(payload)
			if value == "red" {
				stateChannel <- bulb.RED
			} else if value == "blue" {
				stateChannel <- bulb.BLUE
			} else if value == "green" {
				stateChannel <- bulb.GREEN
			}
		})

		client.Subscribe("hermes/Intent/switch_light", func(topic string, payload []byte) {
			if parse(payload) == "on" {
				stateChannel <- bulb.ON
			} else {
				stateChannel <- bulb.OFF
			}
		})

		for {
			bulb.SetState(sc, <-stateChannel)
		}
	})
}

func parse(payload []byte) string {
	var obj map[string]interface{}
	json.Unmarshal(payload, &obj)
	log.Printf("received %v", obj)

	parsed_slots := obj["parsed_slots"].([]interface{})
	slot := parsed_slots[0].(map[string]interface{})
	value := slot["value"].(string)
	log.Printf("value is %v", value)

	return strings.ToLower(value)
}
