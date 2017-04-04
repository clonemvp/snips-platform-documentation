package models

import (
	"encoding/json"
	"log"
)

type Intent struct {
	Text string `json:"text"`
	Intent struct {
		Name        string `json:"intent_name"`
		Probability float64 `json:"probability"`
	} `json:"intent"`

	Slots []struct {
		Name  string `json:"slot_name"`
		Range []int `json:"range"`
		Value string `json:"value"`
	} `json:"slots"`
}

func ParseIntent(payload []byte) Intent {
	var intent Intent
	json.Unmarshal(payload, &intent)
	log.Printf("ParseIntent: received %+v", intent)
	return intent
}
