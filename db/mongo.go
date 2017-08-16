package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

type MongoConnection struct {
	originalSession *mgo.Session
}

func NewDBConnection(mongo *configs.Mongo) (conn *MongoConnection) {
	conn = new(MongoConnection)

	if err := conn.createConnection(mongo); err != nil {
		fmt.Errorf("open error: %s", err)
	}

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

	c.originalSession.SetMode(mgo.Monotonic, true)

	return nil
}

func (c *MongoConnection) CloseConnection() {
	if c.originalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.originalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}

func FillDataBase() {
	connection := NewDBConnection(configs.ConfigInfo.Mongo)
	defer connection.CloseConnection()

	users := connection.GetCollection(configs.ConfigInfo.Mongo)

	for _, user := range FakeUsers {
		err := users.Insert(&user)
		if err != nil {
			fmt.Println("Bad insert")
			break
		}
	}
}
