package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug bool `yml:"debug"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config

	file, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
