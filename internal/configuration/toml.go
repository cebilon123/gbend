package configuration

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// EntitiesMap represents entities. Each entity contains field and type of that field.
type EntitiesMap map[string]map[string]string

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
