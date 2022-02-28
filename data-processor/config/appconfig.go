package config

type StageConfig struct {
	Name      string
	Chunksize int
}
type AppConfig struct {
	Stages map[int]StageConfig
}

var appConfig AppConfig

func SetAppConfig(appConfig AppConfig) {
	appConfig = appConfig
}

func GetAppConfig() AppConfig {
	return appConfig
}
