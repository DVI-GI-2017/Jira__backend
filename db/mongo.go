package db

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Wrapper around mgo.Session
type MongoConnection struct {
	OriginalSession *mgo.Session
}

// Creates new connection to database from config.
func NewDBConnection() (*MongoConnection, error) {
	conn := new(MongoConnection)

	if err := conn.createConnection(configs.ConfigInfo.Mongo); err != nil {
		return conn, fmt.Errorf("open error: %s", err)
	}

	return conn, nil
}

// Drops database associated with this connection if it was set in configuration.
func (c *MongoConnection) DropDataBase(mongo *configs.Mongo) (err error) {
	if mongo.Drop {
		err = c.OriginalSession.DB(mongo.Db).DropDatabase()
		if err != nil {
			return
		}
	}

	return nil
}

// Shortcut for to get database from session.
func (c *MongoConnection) GetDB() (collection *mgo.Database) {
	return c.OriginalSession.DB(configs.ConfigInfo.Mongo.Db)
}

// Returns collection with specified name
func (c *MongoConnection) GetCollection(collection string) *mgo.Collection {
	return c.GetDB().C(collection)
}

// Creates mgo index from custom index type.
func (c *MongoConnection) CreateIndex(collection *mgo.Collection, index *DBIndex) (err error) {
	err = collection.EnsureIndex(mgo.Index{
		Key:        index.Key,
		Unique:     index.Unique,
		DropDups:   index.DropDups,
		Background: index.Background,
		Sparse:     index.Sparse,
	})

	return
}

// Establishes connection to mongo server specified in "mongo" config.
func (c *MongoConnection) createConnection(mongo *configs.Mongo) (err error) {
	fmt.Println("Connecting to local mongo server....")

	c.OriginalSession, err = mgo.Dial(mongo.URL())

	if err != nil {
		return
	}

	c.OriginalSession.SetMode(mgo.Monotonic, true)

	return nil
}

// Insert object into collection
func (c *MongoConnection) Insert(collection string, model interface{}) (result interface{}, err error) {
	if err := c.GetCollection(collection).Insert(model); err != nil {
		return model, err
	}

	return model, nil
}

// Object from collection partially or fully specified by "model"
// Example: collection="users", model = User{Email: "email@mail.ru"} returns
// all users with email == "email@mail.ru"
func (c *MongoConnection) Find(collection string, model interface{}) (result interface{}, err error) {
	result = models.GetModel(tools.GetType(model))

	err = c.GetCollection(collection).Find(model).One(result)

	if err != nil {
		tools.SetParam2Model(result, "IsAuth", false)
	}

	return
}

// Returns all objects from specified collection
func (c *MongoConnection) FindAll(collection string) (result []interface{}, err error) {
	err = c.GetCollection(collection).Find(nil).All(result)
	return
}

// Closes current session.
func (c *MongoConnection) CloseConnection() {
	if c.OriginalSession != nil {
		fmt.Println("Closing local mongo server....")

		c.OriginalSession.Close()

		fmt.Println("Mongo server is closed....")
	}
}

func setFinderInterface(mapModel map[string]string) (finder []interface{}) {
	for key, value := range mapModel {
		finder = append(finder, bson.M{
			key: value,
		})
	}

	return
}
