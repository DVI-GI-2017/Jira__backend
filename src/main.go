package main

import (
	"net/http"
	"router"
	"fmt"
	"strconv"
)

var (
	ctx  *Context
)

func main() {

	var err error

	ctx = app.GetContext()
	defer ctx.Close()

	router.NewRouter(routes)

	port := 3000

	fmt.Printf("Server started on port %d...\n", port)
	http.ListenAndServe(":" + strconv.Itoa(ctx.Config.Port), nil)
}
