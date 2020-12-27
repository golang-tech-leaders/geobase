package main

import (
	"flag"
	"geobase/internal/config"
	"geobase/internal/database"
	"geobase/internal/logger"

	"geobase/internal/server"
)

func main() {
	configFilePath := getConfigFile()
	cfg := config.PrepareConfig(configFilePath)
	logger := logger.New(cfg.LogConf)
	db := database.New()
	db.Init()
	srv := server.NewServer(&cfg.AppConf, db, logger)
	srv.Run()
}

func getConfigFile() string {
	configFile := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	return *configFile
}
