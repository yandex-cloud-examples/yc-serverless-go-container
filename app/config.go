package main

import (
	"github.com/kelseyhightower/envconfig"
)

var c Config

type Config struct {
	Port     int    `envconfig:"PORT"`
	LogLevel string `envconfig:"LOG_LEVEL"`
}

func initConfigFromEnv() {
	envconfig.MustProcess("myapp", &c)
}
