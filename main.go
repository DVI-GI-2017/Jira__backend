package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"strconv"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/db"
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

	mux := startRouter()

	db.FillDataBase()

	fmt.Printf("Server started on port %d...\n", configs.ConfigInfo.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configs.ConfigInfo.Server.Port), mux))
}
