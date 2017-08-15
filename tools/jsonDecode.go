package tools

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"Jira__backend/configs"
)

func decode(r io.Reader) (x *configs.Server, err error) {
	x = new(configs.Server)
	err = json.NewDecoder(r).Decode(x)

	return
}

func GetServerPort(path string) (port int, err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}
	defer file.Close()

	decodeConfig, decodeError := decode(file)

	if decodeError != nil {
		err = decodeError

		return
	}

	port, fileError := strconv.Atoi(decodeConfig.Port)

	err = fileError
	return
}

func GetMongoInfo(path string) {
	// return ServerConfig
}
