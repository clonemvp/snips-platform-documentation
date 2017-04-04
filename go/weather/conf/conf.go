package conf

import (
	"snips-platform-lambda-samples/go/common/conf"
)

type WeatherConf struct {
	PlatformConf conf.PlatformConf
	ApiKey       string
}

func Run(args []string, f func(wc WeatherConf)) {
	conf.Run(args,
		"~/go/bin/weather",
		"Snips Platform sample for weather",
		func(pc conf.PlatformConf) {
			hue := pc.Cfg.Section("OPENWEATHERMAP")
			apiKey := hue.Key("ApiKey").String()

			f(WeatherConf{
				PlatformConf: pc,
				ApiKey:       apiKey,
			})
		})
}
