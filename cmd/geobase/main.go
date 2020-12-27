package main

import (
	"flag"
	"log"

	"geobase/internal/config"
	"geobase/internal/database"
	"geobase/internal/logger"
	"geobase/internal/server"
)

func main() {
	configFilePath := getConfigFile()
	cfg, err := config.PrepareConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(cfg.LogConf)
	urlFinder := database.NewURLFinder()
	locFinder, err := database.NewLocationFinder(cfg.AppConf.DataPath, cfg.AppConf.MetersInRadius)
	if err != nil {
		log.Fatal(err)
	}
	urlFinder.Init()
	srv := server.NewServer(&cfg.AppConf, urlFinder, locFinder, l)
	l.Info().Msg("starting server")
	log.Fatal(srv.Run())
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yml", "config file")
	flag.Parse()
	return *configFile
}
