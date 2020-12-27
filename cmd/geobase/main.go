package main

import (
	"flag"
	"log"

	"geobase/internal/config"
	"geobase/internal/database"
	"geobase/internal/logger"
	"geobase/internal/server"
	"geobase/internal/storage"
)

func main() {
	configFilePath := getConfigFile()
	cfg, err := config.PrepareConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(cfg.LogConf)
	urlFinder := database.New()
	locFinder, err := storage.New(cfg.AppConf.DataPath)
	if err != nil {
		log.Fatal(err)
	}
	urlFinder.Init()
	srv := server.NewServer(&cfg.AppConf, urlFinder, locFinder, l)
	log.Fatal(srv.Run())
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yml", "config file")
	flag.Parse()
	return *configFile
}
