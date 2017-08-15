package main

import (
	"net/http"
	"Jira__backend/router"
	"fmt"
	"strconv"
)

func main() {
	router.NewRouter()

	port := 3000

	fmt.Printf("Server started on port %d...\n", port)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
