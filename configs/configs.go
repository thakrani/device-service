package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port         string `json:"port"`
	Password     string `json:"username"`
	Username     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

var config Config

func LoadConfig(filePath string) (*Config, error) {
	// Open the config file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	// Decode the JSON into the struct
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("could not decode config file: %v", err)
	}

	return &config, nil
}
