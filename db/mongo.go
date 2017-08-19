package db

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	UserCollection    = "users"
	ProjectCollection = "project"
)

type MongoConnection struct {
	OriginalSession *mgo.Session
}

func NewDBConnection() (*MongoConnection, error) {
	conn := new(MongoConnection)

	if err := conn.createConnection(configs.ConfigInfo.Mongo); err != nil {
		return conn, fmt.Errorf("open error: %s", err)
	}

	return conn, nil
}

func (c *MongoConnection) DropDataBase(mongo *configs.Mongo) (err error) {
	if mongo.Drop {
		err = c.OriginalSession.DB(mongo.Db).DropDatabase()
		if err != nil {
			return
		}
	}

	return nil
}

func (c *MongoConnection) GetDB() (collection *mgo.Database) {
	return c.OriginalSession.DB(configs.ConfigInfo.Mongo.Db)
}

func (c *MongoConnection) GetCollection(collectionName string) (collection *mgo.Collection) {
	return c.OriginalSession.DB(configs.ConfigInfo.Mongo.Db).C(collectionName)
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

	c.OriginalSession, err = mgo.Dial(mongo.URL())

	if err != nil {
		return
	}

	c.OriginalSession.SetMode(mgo.Monotonic, true)

	return nil
}

func (c *MongoConnection) Insert(collection string, model interface{}) (result interface{}, err error) {
	if err := c.GetCollection(collection).Insert(model); err != nil {
		return model, err
	}

	return model, nil
}

func (c *MongoConnection) Find(collection string, model interface{}) (result interface{}, err error) {
	result = tools.GetModel(tools.GetType(model))

	var top map[string]string = tools.ParseModel(model)

	fmt.Println("data:")
	fmt.Println(top)

	err = c.GetCollection(collection).Find(bson.M{
		"$and": []interface{}{},
	}).One(&result)

	fmt.Print(result)
	fmt.Printf("\n")

	return
}

func (c *MongoConnection) CloseConnection() {
	if c.OriginalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.OriginalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}

//func FillDataBase() {
//
//	users := Connection.GetCollection(UserCollection)
//
//	for _, user := range FakeUsers {
//		err := users.Insert(&user)
//		if err != nil {
//			fmt.Println("Bad insert")
//			break
//		}
//	}
//}
