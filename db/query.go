package db

import "gopkg.in/mgo.v2"

// Interface for generic queries.
type Query interface {
	All() (result []interface{}, err error)
	One() (result interface{}, err error)
	Count() (count int, err error)
}

// Wrapper around *mgo.Query
type MongoQuery struct {
	*mgo.Query
}

// Get all objects from query.
func (q MongoQuery) All() (result []interface{}, err error) {
	result = make([]interface{}, 0)
	err = q.Query.All(result)
	return
}

// Get one object from query.
func (q MongoQuery) One() (result interface{}, err error) {
	result = new(interface{})
	err = q.Query.One(result)
	return
}

// Count objects in query
func (q MongoQuery) Count() (count int, err error) {
	return q.Query.Count()
}
