package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DatabaseUrl string `json:"databaseUrl"`
	App_Port    string `json:"app_port"`
	User        string `json:"user"`
	DBname      string `json:"dbname"`
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
	// &Config{
	// 	DatabaseUrl: ,
	// 	User: config.User,
	// 	DBname: config.DBname,
	// 	Hostname: config.Hostname,
	// 	Port: config.Port,
	// 	Password: config.Password,
	// 	Ssl: config.Ssl,
	// },
	return config, nil
}
