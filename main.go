package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/login"
)

func main() {
	login.InitKeys()

	config, err := configs.FromFile("config.json")

	if err != nil {
		log.Panic("bad configs: ", err)
	}

	db.NewDBConnection(config.Mongo)

	mux, err := routes.NewRouter()

	if err != nil {
		log.Panic("can not create router: ", err)
	}

	fmt.Printf("Server started on port %d...\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), mux))
}
