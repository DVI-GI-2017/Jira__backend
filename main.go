package main

import (
	"log"
	"os/exec"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/mux"
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

func init() {
	_, err := exec.Command("sh", "-c",
		"mkdir rsa && cd rsa && "+
			"openssl genrsa -out app.rsa 1024 && "+
			"openssl rsa -in app.rsa -pubout > app.rsa.pub").Output()
	if err != nil {
		log.Println(err)
	}
	pool.InitWorkers()
	rsaInit()
}

func main() {
	config, err := configs.FromFile("config.json")
	if err != nil {
		log.Panicf("can not init config: %v", err)
	}

	db.InitDB(config.Mongo)

	router, err := mux.NewRouter("/api/v1")
	if err != nil {
		log.Fatalf("can not create router: %v", err)
	}
	router.AddWrappers(mux.Logger)

	routes.SetupRoutes(router)

	router.PrintRoutes()

	port := config.Server.GetPort()

	log.Printf("Server started on port %s...\n", port)

	log.Fatal(router.ListenAndServe(port))
}
