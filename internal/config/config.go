package config

import (
	"gandiwa/pkg/viper"
	"log"
)

type Config struct {
	AppPort string
}

func LoadConfigFile(configPath string) *Config {
	conf := Config{}

	if err := viper.LoadYAMLToStruct(configPath, &conf); err != nil {
		log.Fatalf("failed to load config. err: %s", err.Error())
	}

	return &conf
}
