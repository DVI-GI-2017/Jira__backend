package main

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func rsaInit() {
	err := auth.InitKeys()

	if err != nil {
		log.Panic("can not init rsa keys: ", err)
	}
}

func startRouter() (mux http.Handler) {
	mux, err := routes.NewRouter()

	if err != nil {
		log.Panic("can not create router: ", err)
	}

	return
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

func main() {
	rsaInit()

	configs.ParseFromFile("config.json")

	db.StartDB()

	raii(db.Connection.CloseConnection)
	db.Connection.DropDataBase(configs.ConfigInfo.Mongo)
	db.FillDataBase()

	mux := startRouter()

	fmt.Printf("Server started on port %d...\n", configs.ConfigInfo.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configs.ConfigInfo.Server.Port), mux))
}
