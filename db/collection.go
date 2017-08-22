package db

import "gopkg.in/mgo.v2"

// Interface for generic collection.
type Collection interface {
	FindId(id interface{}) Query
	Find(selector interface{}) Query
}

// Wrapper around *mgo.Collection
type MongoCollection struct {
	*mgo.Collection
}

// Return Query with objects with given id
func (c MongoCollection) FindId(id interface{}) Query {
	return MongoQuery{c.Collection.FindId(id)}
}

// Return Query with objects that match given query
func (c MongoCollection) Find(query interface{}) Query {
	return MongoQuery{c.Collection.Find(query)}
}
