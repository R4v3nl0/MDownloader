package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(filepath string) (*Config, error) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		return &Config{}, err
	}

	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return &Config{}, err
	}

	return config, nil
}
