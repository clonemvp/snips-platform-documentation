package bulb

import (
	"encoding/json"
	"fmt"
	"log"
	"bytes"
	"net/http"
	"snips-platform-lambda-samples/go/hue/conf"
)

type State struct {
	On  bool `json:"on"`
	Hue int `json:"hue"`
	Sat int `json:"sat"`
	Bri int `json:"bri"`
}

func newOnState(hue int, sat int) State {
	return State{
		On:  true,
		Hue: hue,
		Sat: sat,
		Bri: 255,
	}
}

func newOffState() State {
	return State{
		On: false,
	}
}

var (
	ON    = newOnState(12750, 128)
	OFF   = newOffState()
	BLUE  = newOnState(46920, 255)
	RED   = newOnState(0, 255)
	GREEN = newOnState(25500, 255)
)

func SetState(sc conf.SampleConf, state State) {
	r, _ := json.Marshal(state)

	for _, bulbId := range sc.HueBulbs {
		url := fmt.Sprintf("http://%v/api/%v/lights/%v/state", sc.HueRouter, sc.HueUsername, bulbId)
		log.Printf("%v on %v", string(r), url)
		request, _ := http.NewRequest("PUT", url, bytes.NewReader(r))
		http.DefaultClient.Do(request)
	}
}
