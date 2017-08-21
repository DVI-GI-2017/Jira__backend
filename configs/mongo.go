package configs

import "fmt"

type Mongo struct {
	Server string `json:"server"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DB     string `json:"db"`
	Drop   bool   `json:"drop"`
}

func (m *Mongo) URL() string {
	return fmt.Sprintf("%s://%s:%d", m.Server, m.Host, m.Port)
}
