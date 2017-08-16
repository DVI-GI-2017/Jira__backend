package configs

import (
	"io"
	"os"
	"fmt"
	"encoding/json"
	"log"
)

type Config struct {
	Mongo  *Mongo
	Server *Server
}

var ConfigInfo *Config

func FromFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return &Config{}, fmt.Errorf("can not open file: %v", err)
	}

	return FromReader(file)
}

func FromReader(r io.Reader) (*Config, error) {
	config := new(Config)
	err := json.NewDecoder(r).Decode(config)

	if err != nil {
		return config, fmt.Errorf("can not parse config: %v", err)
	}

	return config, nil
}

func ParseFromFile(path string) {
	config, err := FromFile(path)

	if err != nil {
		log.Panic("bad configs: ", err)
	}

	ConfigInfo = config
}
