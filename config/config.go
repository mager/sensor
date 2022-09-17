package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	OpenWeatherMapKey string
}

func ProvideConfig() Config {
	var cfg Config
	err := envconfig.Process("sensor", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}

var Options = ProvideConfig
