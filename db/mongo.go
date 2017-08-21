package db

import (
	"fmt"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"gopkg.in/mgo.v2"
)

// Wrapper around mgo.Session
var database *mgo.Database

// Initialize global database instance
func InitDB(mongo *configs.Mongo) error {
	log.Println("Connecting to local mongo server....")

	session, err := mgo.Dial(mongo.URL())
	if err != nil {
		return fmt.Errorf("can not connect to database: %v", err)
	}

	session.SetMode(mgo.Monotonic, true)

	database = session.DB(mongo.DB)

	return nil
}

// Returns current db with new session
func GetDB() *mgo.Database {
	return database.With(database.Session.Copy())
}
