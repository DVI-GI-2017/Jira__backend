package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

func rsaInit() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}
}

func initPort() (port string) {
	port = os.Getenv("PORT")

	if port == "" {
		port = strconv.Itoa(configs.ConfigInfo.Server.Port)
	}

	return
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

	routes.InitRouter(router, routes.LoginRoutes)
	routes.InitRouter(router, routes.TestRoutes)
	routes.InitRouter(router, routes.UsersRoutes)

	port := initPort()

	fmt.Printf("Server started on port %s...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.Logger(router)))
}
