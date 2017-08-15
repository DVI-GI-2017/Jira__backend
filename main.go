package main

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
	"Jira__backend/tools"
	"Jira__backend/router"
)

func main() {
	port, err := tools.GetPort("configs/server.json")

	if err != nil {
		panic("bad config")
	}

	router.NewRouter()

	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
