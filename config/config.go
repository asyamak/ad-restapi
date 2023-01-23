package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
	User string `json:"user"`
	DBname string `json:"dbname"`
	Hostname string `json:"hostname"`
	Port string `json:"port"`
	Password string `json:"postgres"`
	Ssl string `json:"ssl"`

}

// const filename = "./config/config.json"

func New(filename string) (*Config, error){
	var config *Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os open: %w",err)
	}

	if err = json.NewDecoder(configFile).Decode(&config); err != nil{
		return nil, fmt.Errorf("decoding: %w",err)
	}

	return &Config{
		User: config.User,
		DBname: config.DBname,
		Hostname: config.Hostname,
		Port: config.Port,
		Password: config.Password,
		Ssl: config.Ssl,
	},nil
}