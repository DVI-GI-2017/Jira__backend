package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/routes"
)

func rsaInit() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}
}

func raii(handler func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		handler()
		os.Exit(0)
	}()
}

func init() {
	rsaInit()

	configs.ParseFromFile("config.json")

	db.StartDB()

	raii(db.Connection.CloseConnection)
	db.Connection.DropDataBase(configs.ConfigInfo.Mongo)
	db.FillDataBase()
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
