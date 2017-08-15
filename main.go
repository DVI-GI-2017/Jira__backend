package main

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
	"Jira__backend/tools"
	"Jira__backend/router"
	"Jira__backend/dataBase"
)

func main() {
	dataBase.NewDBConnection()

	port, err := tools.GetServerPort("configs/server.json")

	if err != nil {
		panic("bad config")
	}

	router.NewRouter()

	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
