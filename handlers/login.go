package handlers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/configs"
)

var RegisterUser = PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var credentials auth.Credentials

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Fprint(w, "Invalid registration data format.")
			log.Printf("%v", err)

			return
		}

		if err := auth.RegisterUser(&credentials); err != nil {
			w.WriteHeader(http.StatusConflict)

			fmt.Fprint(w, "User with this email already exists.")
			log.Printf("%v", err)

			return
		}

		user := models.User{Email: credentials.Email, Password: credentials.Password}

		db.FakeUsers = append(db.FakeUsers, user)
		w.WriteHeader(http.StatusOK)

		connection := db.NewDBConnection(configs.ConfigInfo.Mongo)
		defer connection.CloseConnection()

		users := connection.GetCollection(configs.ConfigInfo.Mongo)

		result := models.User{}
		err := users.Find(bson.M{"firstname": "Jeremy"}).Select(bson.M{"Email": 0}).One(&result)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprint(w, "Bad find")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, result)
	})

var Login = PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var user auth.Credentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, "Error in request")
			log.Printf("%v", err)

			return
		}

		if err := auth.LoginUser(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, err)
			log.Printf("%v", err)

			return
		}

		token, err := auth.NewToken()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("%v", err)

			return
		}

		response := token
		tools.JsonResponse(response, w)
	})
