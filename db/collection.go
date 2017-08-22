package db

import "gopkg.in/mgo.v2"

// Interface for generic collection.
type Collection interface {
	FindId(id interface{}) Query
}

// Wrapper around *mgo.Collection
type MongoCollection struct {
	*mgo.Collection
}

// Return Query with objects
func (c MongoCollection) FindId(id interface{}) Query {
	return MongoQuery{c.Collection.FindId(id)}
}
