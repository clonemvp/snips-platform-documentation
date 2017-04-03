package models

import (
	"encoding/json"
	"log"
)

type Intent struct {
	Text string `json:"text"`
	ParsedIntent struct {
		IntentName  string `json:"intent_name"`
		Probability float64
	} `json:"parsed_intent"`

	ParsedSlots []struct {
		Value      string `json:"value"`
		MatchRange []int `json:"match_range"`
		SlotName   string `json:"slot_name"`
	} `json:"parsed_slots"`
}

func ParseIntent(payload []byte) Intent {
	var intent Intent
	json.Unmarshal(payload, &intent)
	log.Printf("ParseIntent: received %+v", intent)
	return intent
}
