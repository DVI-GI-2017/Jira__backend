package configs

import "fmt"

type Mongo struct {
	Server      string `json:"server"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Db          string `json:"db"`
	Drop        bool   `json:"drop"`
	Connections int    `json:"connections"`
}

func (m *Mongo) URL() string {
	return fmt.Sprintf("%s://%s:%d", m.Server, m.Host, m.Port)
}
