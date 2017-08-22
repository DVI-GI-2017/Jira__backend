package db

import (
	"fmt"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/configs"
	"gopkg.in/mgo.v2"
)

// Wrapper around *mgo.Session.
type MongoSession struct {
	*mgo.Session
}

// Override DB method of mgo.Session to return wrapper around *mgo.DataSource.
func (s MongoSession) DB(name string) DataSource {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// Wrapper around *mgo.DataSource.
type MongoDatabase struct {
	*mgo.Database
}

// Override C method of mgo.DataSource to return wrapper around *mgo.Collection
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// Wrapper around *mgo.Collection
type MongoCollection struct {
	*mgo.Collection
}

// Default defaultDB
var defaultDB DataSource

// Initialize global defaultDB instance
func InitDB(mongo *configs.Mongo) {
	log.Println("Connecting to local mongo server....")

	session, err := NewMongoSession(mongo.URL())
	if err != nil {
		log.Panicf("can not connect to mongo server: %v", err)
	}

	defaultDB = session.DB(mongo.DB)
}

// Creates new mongo session
func NewMongoSession(mgoURI string) (Session, error) {
	mgoSession, err := mgo.Dial(mgoURI)
	if err != nil {
		return nil, fmt.Errorf("can not open defaultDB session: %v", err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return MongoSession{mgoSession}, nil
}

// Returns current data source with new session
func Copy() DataSource {
	return defaultDB.With(defaultDB.Session.Copy())
}

// Returns current data source.
func Get() DataSource {
	return defaultDB
}
