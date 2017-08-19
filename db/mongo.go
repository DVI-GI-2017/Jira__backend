package db

import (
	"fmt"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2"
)

const (
	UserCollection    = "users"
	ProjectCollection = "project"
)

type MongoConnection struct {
	originalSession *mgo.Session
}

var Connection *MongoConnection

func NewDBConnection(mongo *configs.Mongo) (*MongoConnection, error) {
	conn := new(MongoConnection)

	if err := conn.createConnection(mongo); err != nil {
		return conn, fmt.Errorf("open error: %s", err)
	}

	return conn, nil
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

func (c *MongoConnection) GetDB() (collection *mgo.Database) {
	return c.originalSession.DB(configs.ConfigInfo.Mongo.Db)
}

func (c *MongoConnection) GetCollection(collectionName string) (collection *mgo.Collection) {
	return c.originalSession.DB(configs.ConfigInfo.Mongo.Db).C(collectionName)
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

	c.originalSession.SetMode(mgo.Monotonic, true)

	return nil
}

func (c *MongoConnection) Insert(collection string) {
	if c.originalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.originalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}

func (c *MongoConnection) CloseConnection() {
	if c.originalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.originalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}

func StartDB() {
	newConnection, err := NewDBConnection(configs.ConfigInfo.Mongo)
	if err != nil || newConnection == nil {
		log.Panicf("can not start db: %s", err)
	}
	Connection = newConnection
}

func FillDataBase() {
	users := Connection.GetCollection(configs.ConfigInfo.Mongo.Db, UserCollection)

	for _, user := range FakeUsers {
		err := users.Insert(&user)
		if err != nil {
			fmt.Println("Bad insert")
			break
		}
	}
}
