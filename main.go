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
	db.NewDBConnection()

	conf, err := configs.FromFile("config.json")

	if err != nil {
		log.Panic("bad configs: ", err)
	}

	port := conf.Server.Port

	router.NewRouter()

	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
