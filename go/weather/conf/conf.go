package conf

import (
	"github.com/go-ini/ini"
	"github.com/jawher/mow.cli"
	"log"
)

type WeatherConf struct {
	Cfg          ini.File
	PlatformHost string
	PlatformPort int
	ApiKey       string
}

func Run(args []string, f func(hc WeatherConf)) {
	app := cli.App("~/go/bin/weather", "Snips Platform sample for weather")
	path := app.StringOpt("c", "conf.ini", "Configuration file")
	app.Action = func() {
		cfg, err := ini.Load(*path)
		if err != nil {
			log.Fatalln("ini.Load: %v", err)
		}

		networkConf := cfg.Section("SNIPS-PLATFORM")
		platformHost := networkConf.Key("Host").String()
		platformPort, _ := networkConf.Key("Port").Int()

		hue := cfg.Section("OPENWEATHERMAP")
		apiKey := hue.Key("ApiKey").String()

		f(WeatherConf{
			Cfg:          *cfg,
			PlatformHost: platformHost,
			PlatformPort: platformPort,
			ApiKey:       apiKey,
		})
	}
	app.Run(args)
}
