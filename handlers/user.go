package handlers

import (
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/models"
)

var CheckDB = GetOnly(
	func(w http.ResponseWriter, r *http.Request) {
		connection := db.NewDBConnection(configs.ConfigInfo.Mongo)
		defer connection.CloseConnection()

		users := connection.GetCollection(configs.ConfigInfo.Mongo)

		err := users.Insert(&db.FakeUsers[0])
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, "Bad insert")
		}

		result := models.User{}
		err = users.Find(bson.M{"firstname": "Jeremy"}).Select(bson.M{"Email": 0}).One(&result)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprint(w, "Bad find")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, result)
	})
