package main

import (
	"flag"

	"github.com/raintank/fakemetrics_ng/agents"
)

var (
	confFile = flag.String(
		"config", "fakemetrics.ini", "configuration file path",
	)
)

func main() {
	flag.Parse()

	setupConfig()

	a := agents.New(timerMod, dataGenMod, outMod)
	a.Run()
}
