package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/router"
	"github.com/DVI-GI-2017/Jira__backend/db"
)

func main() {
	db.NewDBConnection()

	port, err := tools.GetServerPort("configs/server.json")

	if err != nil {
		panic("bad config")
	}

	router.NewRouter()

	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
