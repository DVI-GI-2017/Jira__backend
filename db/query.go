package db

import "gopkg.in/mgo.v2"

// Interface for generic queries.
type Query interface {
	// Returns all query entries
	// NOTE: We use just interface{} not []interface{} because
	// conversion from interface{} to Type is O(1) but conversion
	// from []interface{} to []Type is O(n)
	All() (result interface{}, err error)

	// Return first object in query
	One() (result interface{}, err error)

	// Count all objects in query
	Count() (count int, err error)

	// Returns true if query is empty
	IsEmpty() (empty bool, err error)
}

// Wrapper around *mgo.Query
type MongoQuery struct {
	*mgo.Query
}

// Get all objects from query.
func (q MongoQuery) All() (result interface{}, err error) {
	err = q.Query.All(result)
	return
}

// Get one object from query.
func (q MongoQuery) One() (result interface{}, err error) {
	err = q.Query.One(result)
	return
}

// Count objects in query
func (q MongoQuery) Count() (count int, err error) {
	return q.Query.Count()
}

// Returns true if query is empty
func (q MongoQuery) IsEmpty() (empty bool, err error) {
	count, err := q.Count()
	return count == 0, err
}
