package configuration

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type EntitiesMap map[string]map[string]any

type Config struct {
	Entities EntitiesMap
}

func LoadConfiguration(fileName string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(fileName, &config)
	if err != nil {
		return nil, fmt.Errorf("decode file: %w", err)
	}

	return &config, nil
}
