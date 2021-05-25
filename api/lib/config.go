package lib

import (
	"github.com/caarlos0/env"
)

type config struct { // nolint:maligned
	Version   string `env:"APP_VERSION" envDefault:""`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"warn"`
	AccessLog bool   `env:"ACCESS_LOG" envDefault:"true"`
}

var (
	ver    = "dev"
	commit string
	date   string
)

// Config represents its configurations
var Config *config

func init() {
	Set()
}

// Set sets configurations via envoronment variables
func Set() {
	Config = &config{}
	env.Parse(Config)
}
