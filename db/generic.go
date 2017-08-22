package db

// Interface for generic session with data source.
type Session interface {
	DB(name string) DataSource
	Close()
}

// Interface for generic data source (e.g. database).
type DataSource interface {
	C(name string) Collection
}

// Interface for generic collection.
type Collection interface {
}
