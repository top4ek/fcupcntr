package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type settings struct {
	Port string `env:"PORT" env-default:"8080"`
	DSN  string `env:"DSN" env-default:"fcupcntr.db"`
}

var Settings settings

func Load() error {
	return cleanenv.ReadEnv(&Settings)
}
