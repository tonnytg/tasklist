package config

import (
	"encoding/json"
	"os"
)

// Create config yaml reader
type Config struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Database    struct {
		Type  string `json:"type"`
		Table string `json:"table"`
		Name  string `json:"name"`
		Path  string `json:"path"`
	} `json:"database"`
}

func Reader() *Config {
	// Read config file
	c := Config{}

	// read configfile
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal json
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	// Return config
	return &c
}
