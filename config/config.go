package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	App      AppConfig      `json:"app_config"`
	DataBase DataBaseConfig `json:"db_config"`
}
type AppConfig struct {
	Port string `json:"app_port"`
}

type DataBaseConfig struct {
	DatabaseURL string `json:"database_url"`
	User        string `json:"user"`
	DbName      string `json:"db_name"`
	Hostname    string `json:"hostname"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	Ssl         string `json:"ssl"`
}

func New(filename string) (*Config, error) {
	var config *Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os open: %w", err)
	}

	if err = json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return config, nil
}
