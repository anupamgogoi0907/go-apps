package main

import (
	"fmt"
	"github.com/anupamgogoi0907/go-apps/data-processor/config"
	"github.com/anupamgogoi0907/go-apps/data-processor/pkg/pipeline"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Starting pipeline...")
	start := time.Now()

	appConfig := loadConfig()
	config.SetAppConfig(appConfig)

	args := os.Args[1:]
	p, err := pipeline.NewPipeline(args...)
	if err == nil {
		p.RunPipeline()
	}
	elapsed := time.Since(start)
	log.Printf("Finishing pipeline.Time taken:%s\n", elapsed)
}

func loadConfig() (appConfig config.AppConfig) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading appConfig file, %s", err)
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return appConfig
}
