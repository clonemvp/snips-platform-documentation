package conf

import (
	"snips-platform-lambda-samples/go/common/conf"
)

type HueConf struct {
	PlatformConf conf.PlatformConf
	HueRouter    string
	HueUsername  string
	HueBulbs     []string
}

func Run(args []string, f func(hc HueConf)) {
	conf.Run(args,
		"~/go/bin/hue",
		"Snips Platform sample using Philips Hue",
		func(pc conf.PlatformConf) {
			hue := pc.Cfg.Section("HUE")
			hueRouter := hue.Key("Router").String()
			hueUsername := hue.Key("Username").String()
			hueBulbs := hue.Key("Bulbs").Strings(",")

			f(HueConf{
				PlatformConf: pc,
				HueRouter:    hueRouter,
				HueUsername:  hueUsername,
				HueBulbs:     hueBulbs,
			})
		})
}
