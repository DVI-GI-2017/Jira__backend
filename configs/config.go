package configs

import (
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	Mongo  *Mongo
	Server *Server
}

func FromFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return &Config{}, fmt.Errorf("can not open file: %v", err)
	}

	config := new(Config)
	err = json.NewDecoder(file).Decode(config)

	if err != nil {
		return config, fmt.Errorf("can not parse json: %v", err)
	}

	return config, nil
}
