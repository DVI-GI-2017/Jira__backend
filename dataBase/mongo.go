package dataBase

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

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
		Key:        []string{"name", "phone"},
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
	err = collection.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	// Query One
	result := Person{}
	err = collection.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)

	// Query All
	var results []Person
	err = collection.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = collection.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = collection.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

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