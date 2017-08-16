package db

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

type MongoConnection struct {
	originalSession *mgo.Session
}

func NewDBConnection(mongo *configs.Mongo) (conn *MongoConnection) {
	conn = new(MongoConnection)
	conn.createConnection(mongo)
	return
}

func (c *MongoConnection) DropDataBase(mongo *configs.Mongo) (err error) {
	if mongo.Drop {
		err = c.originalSession.DB(mongo.Db).DropDatabase()
		if err != nil {
			return
		}
	}

	return nil
}

func (c *MongoConnection) GetCollection(mongo *configs.Mongo) (collection *mgo.Collection) {
	return c.originalSession.DB(mongo.Db).C(mongo.Collections[0])
}

func (c *MongoConnection) SetIndex(collection *mgo.Collection, index *tools.DBIndex) (err error) {
	err = collection.EnsureIndex(mgo.Index{
		Key:        index.Key,
		Unique:     index.Unique,
		DropDups:   index.DropDups,
		Background: index.Background,
		Sparse:     index.Sparse,
	})

	return
}

func (c *MongoConnection) createConnection(mongo *configs.Mongo) (err error) {
	fmt.Println("Connecting to local mongo server....")

	c.originalSession, err = mgo.Dial(mongo.URL())

	if err != nil {
		return
	}

	defer c.CloseConnection()

	c.originalSession.SetMode(mgo.Monotonic, true)

	// TODO: Init several collections or remove they from config?
	users := c.originalSession.DB(mongo.Db).C(mongo.Collections[0])

	err = c.SetIndex(users, &tools.DBIndex{
		Key:        []string{"first_name", "updated_at"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	if err != nil {
		return err
	}

	// Insert Datas
	err = users.Insert(&FakeUsers[0])

	if err != nil {
		panic(err)
	}

	// Query One
	result := models.User{}
	err = users.Find(bson.M{"firstname": "Jeremy"}).Select(bson.M{"Email": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email", result)

	// Query All
	var results models.Users
	err = users.Find(bson.M{"firstname": "Jeremy"}).Sort("-created_at").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Jeremy"}
	change := bson.M{"$set": bson.M{"lastname": "Cumberbatch", "updated_at": time.Now()}}
	err = users.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = users.Find(bson.M{"firstname": "Jeremy"}).Sort("-updated_at").All(&results)

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
