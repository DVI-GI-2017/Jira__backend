package models

type MongoConfig struct {
	Server string `json:"server"`
	Host string `json:"host"`
	Port string `json:"port"`
	Db string `json:"db"`
	Collections []string `json:"collections"`
}
