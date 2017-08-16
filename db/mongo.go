package db

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/models"
)

var (
	IsDrop = true
)

type MongoConnection struct {
	originalSession *mgo.Session
}

func NewDBConnection() (conn *MongoConnection) {
	conn = new(MongoConnection)
	conn.createConnection()
	return
}

func (c *MongoConnection) createConnection() (err error) {
	fmt.Println("Connecting to local mongo server....")
	c.originalSession, err = mgo.Dial("mongodb://127.0.0.1:27017")

	if err != nil {
		return
	}

	defer c.CloseConnection()

	c.originalSession.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = c.originalSession.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	collection := c.originalSession.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key:        []string{"first_name", "updated_at"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = collection.Insert(&FakeUsers[0])

	if err != nil {
		panic(err)
	}

	// Query One
	result := models.User{}
	err = collection.Find(bson.M{"first_name": "Jeremy"}).Select(bson.M{"Email": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email", result)

	// Query All
	var results models.Users
	err = collection.Find(bson.M{"first_name": "Jeremy"}).Sort("-created_at").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Jeremy"}
	change := bson.M{"$set": bson.M{"last_name": "Cumberbatch", "updated_at": time.Now()}}
	err = collection.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = collection.Find(bson.M{"first_name": "Jeremy"}).Sort("-updated_at").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	return nil
}

func (c *MongoConnection) CloseConnection() {
	if c.originalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.originalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}
