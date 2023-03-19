package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	DBConnectionStr string `json:"dbConnectionStr"`
}

func Load() (AppConfig, error) {
	var env = os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	data, err := os.ReadFile("config." + env + ".json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return AppConfig{}, err
	}

	config := AppConfig{}
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return AppConfig{}, err
	}

	return config, nil
}

var instance AppConfig //nolint: gochecknoglobals // Use singleton pattern

func GetConfig() AppConfig {
	var once sync.Once
	once.Do(func() {
		config, err := Load()
		if err == nil {
			instance = config
		}
	})
	return instance
}
