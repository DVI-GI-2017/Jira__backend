package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"os/signal"
	"syscall"
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

func configParse(path string) (config *configs.Config) {
	config, err := configs.FromFile(path)

	if err != nil {
		log.Panic("bad configs: ", err)
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

func startRouter() (mux http.Handler) {
	mux, err := routes.NewRouter()

	if err != nil {
		log.Panic("can not create router: ", err)
	}

	return
}

func main() {
	rsaInit()

	config := configParse("config.json")
	connection := db.NewDBConnection(config.Mongo)

	raii(connection.CloseConnection)
	mux := startRouter()

	fmt.Printf("Server started on port %d...\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), mux))
}
