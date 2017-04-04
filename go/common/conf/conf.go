package conf

import (
	"github.com/go-ini/ini"
	"github.com/jawher/mow.cli"
	"log"
)

type PlatformConf struct {
	Cfg          ini.File
	PlatformHost string
	PlatformPort int
}

func Run(args []string, appName string, appDesc string, f func(PlatformConf)) {
	app := cli.App(appName, appDesc)
	path := app.StringOpt("c", "conf.ini", "Configuration file")
	app.Action = func() {
		cfg, err := ini.Load(*path)
		if err != nil {
			log.Fatalln("ini.Load: %v", err)
		}

		networkConf := cfg.Section("SNIPS-PLATFORM")
		platformHost := networkConf.Key("Host").String()
		platformPort, _ := networkConf.Key("Port").Int()

		f(PlatformConf{
			Cfg:          *cfg,
			PlatformHost: platformHost,
			PlatformPort: platformPort,
		})
	}
	app.Run(args)
}
