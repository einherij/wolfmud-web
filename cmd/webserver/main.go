package main

import (
	"fmt"

	"code.wolfmud.org/WolfMUD.git/comms"
	"code.wolfmud.org/WolfMUD.git/config"
	"code.wolfmud.org/WolfMUD.git/stats"
	"code.wolfmud.org/WolfMUD.git/zones"

	webconfig "github.com/einherij/wolfmud-web/pkg/config"
	"github.com/einherij/wolfmud-web/pkg/server"
)

func main() {
	go RunTCPServer()
	RunWebServer()
}

func RunTCPServer() {
	stats.Start()
	zones.Load()
	comms.Listen(config.Server.Host, config.Server.Port)
}

func RunWebServer() {
	conf, err := webconfig.ParseConfig()
	if err != nil {
		panic(fmt.Errorf("error parsing web server config: %w", err))
	}
	webSrv := server.New(conf)
	webSrv.Run()
}
