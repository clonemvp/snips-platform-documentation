package conf

import (
	"github.com/go-ini/ini"
	"github.com/jawher/mow.cli"
	"log"
)

type SampleConf struct {
	Cfg          ini.File
	PlatformHost string
	PlatformPort int
	HueRouter    string
	HueUsername  string
	HueBulbs     []string
}

func Run(args []string, f func(hc SampleConf)) {
	app := cli.App("~/go/bin/hue", "Snips Platform sample using Philips Hue")
	path := app.StringOpt("c", "conf.ini", "Configuration file")
	app.Action = func() {
		cfg, err := ini.Load(*path)
		if err != nil {
			log.Fatalln("ini.Load: %v", err)
		}

		networkConf := cfg.Section("SNIPS-PLATFORM")
		platformHost := networkConf.Key("Host").String()
		platformPort, _ := networkConf.Key("Port").Int()

		hue := cfg.Section("HUE")
		hueRouter := hue.Key("Router").String()
		hueUsername := hue.Key("Username").String()
		hueBulbs := hue.Key("Bulbs").Strings(",")

		f(SampleConf{
			Cfg:          *cfg,
			PlatformHost: platformHost,
			PlatformPort: platformPort,
			HueRouter:    hueRouter,
			HueUsername:  hueUsername,
			HueBulbs:     hueBulbs,
		})
	}
	app.Run(args)
}
