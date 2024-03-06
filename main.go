package main

import (
	"eintrag/pkg/eintrag"
	"flag"
)

func main() {
	configFile := flag.String("config", "config.json", "")
	flag.Parse()

	app := eintrag.NewApp(configFile)

	app.Init()
	app.Run()
}
