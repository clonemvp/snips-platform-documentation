package main

import (
	"os"
	"strings"
	"snips-platform-lambda-samples/go/common/mqtt"
	"snips-platform-lambda-samples/go/hue/conf"
	"snips-platform-lambda-samples/go/hue/bulb"
	"snips-platform-lambda-samples/go/common/models"
)

func main() {
	conf.Run(os.Args, func(hc conf.HueConf) {
		client := mqtt.NewClient(hc.PlatformHost, hc.PlatformPort)
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
			bulb.SetState(hc, <-stateChannel)
		}
	})
}

func parse(payload []byte) string {
	intent := models.ParseIntent(payload)
	value := intent.ParsedSlots[0].Value
	return strings.ToLower(value)
}
