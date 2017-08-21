package db

type DBIndex struct {
	Key        []string
	Unique     bool
	DropDups   bool
	Background bool
	Sparse     bool
}
