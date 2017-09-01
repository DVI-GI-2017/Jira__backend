package main

import (
	"log"

	_ "github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/weitbelou/yac"
)

func main() {
	config, err := configs.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	db.InitDB(config.Mongo)

	router, err := yac.NewRouter("/api/v1")
	if err != nil {
		log.Fatalf("can not create router: %v", err)
	}
	router.AddWrappers(yac.Logger)

	routes.SetupRoutes(router)

	router.PrintRoutes()

	port := config.Server.GetPort()

	log.Printf("Server started on port %s...\n", port)

	log.Fatal(router.ListenAndServe(port))
}
