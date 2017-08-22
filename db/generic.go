package db

// Interface for generic session with data source.
type Session interface {
	// Returns data source associated with this session with "name"
	Source(name string) DataSource

	// Closes session with data source other methods should panic
	// if you are trying to use closed session
	Close()
}

// Interface for generic data source (e.g. database).
type DataSource interface {
	// Returns collection by name
	C(name string) Collection

	// Returns copy of data source (may be copy of session as well)
	Copy() DataSource
}

// Interface for generic collection.
type Collection interface {
}
