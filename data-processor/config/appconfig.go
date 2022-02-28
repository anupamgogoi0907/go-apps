package config

import (
	"log"
)

type StageConfig struct {
	Name      string
	Chunksize int
	Noworkers int
}
type AppConfig struct {
	Stages map[int]StageConfig
}

var appConfig AppConfig

func SetAppConfig(cfg AppConfig) {
	appConfig = cfg
	log.Println("Config file loaded.")
}

func GetAppConfig() AppConfig {
	return appConfig
}
