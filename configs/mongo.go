package configs

type Mongo struct {
	Server string `json:"server"`
	Host string `json:"host"`
	Port int `json:"port"`
	Db string `json:"db"`
	Collections []string `json:"collections"`
}
