package config

import (
	"encoding/json"
	"os"
)

// Create config json reader
type Config struct {
	Name string
	Path string
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
