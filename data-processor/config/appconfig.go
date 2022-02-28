package config

import (
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/utility"
	"log"
)

type StageConfig struct {
	Name            string
	Chunksizestring string
	Chunksize       int
	Noworkers       int
}
type AppConfig struct {
	Stages map[int]StageConfig
}

var appConfig AppConfig

func SetAppConfig(cfg AppConfig) {
	m := cfg.Stages[1]
	m.Chunksize = utility.GetChunkSize(m.Chunksizestring)
	cfg.Stages[1] = m
	appConfig = cfg
	log.Println("Config file loaded.")
}

func GetAppConfig() AppConfig {
	return appConfig
}
