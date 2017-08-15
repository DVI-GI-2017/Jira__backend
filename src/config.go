package main

import (
	"os"
	"path/filepath"
	"io/ioutil"

	"encoding/json"
)

type config struct {
	Mongo mongoConf `json:"Mongo"`
}

type mongoConf struct {
	Address  string `json:"Address"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
	Timeout  uint   `json:"Timeout"`
}

func loadConfig() (conf *config, err error) {
	var (
		file, dir string
		bytes     []byte
	)

	if file = os.Getenv("CONFIG"); file == "" {
		if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			return
		}
		file = dir + "config.json"
	}

	if bytes, err = ioutil.ReadFile(file); err != nil {
		return
	}

	conf = new(config)
	if err = json.Unmarshal(bytes, conf); err != nil {
		return
	}

	if muerr := conf.Validate(); muerr.HasErrors() {
		err = muerr
	}

	return
}