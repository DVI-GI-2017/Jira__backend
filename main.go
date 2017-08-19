package main

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"log"
	"net/http"
)

func rsaInit() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}
}

func init() {
	configs.ParseFromFile("config.json")

	pool.InitWorkers()
	rsaInit()
}

func main() {
	router, err := routes.NewRouter("/api/v1")
	if err != nil {
		log.Fatalf("can not create router: %v", err)
	}
	routes.InitRouter(router)

	port := configs.ConfigInfo.Server.Port

	fmt.Printf("Server started on port %d...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.Logger(router)))
}
