package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/DVI-GI-2017/Jira__backend/router"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/configs"
)

func main() {
	config, err := configs.FromFile("config.json")

	if err != nil {
		log.Panic("bad configs: ", err)
	}

	db.NewDBConnection(config.Mongo)

	router.NewRouter()

	fmt.Printf("Server started on port %d...\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), nil))
}
