package handlers

import (
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"encoding/json"
	"fmt"
	"log"
)

var AllUsers = GetOnly(
	func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(db.FakeUsers)

		if err != nil {
			fmt.Fprint(w, "Something happened...")
			log.Printf("%v", err)

			w.WriteHeader(http.StatusBadRequest)
		}
	},
)
