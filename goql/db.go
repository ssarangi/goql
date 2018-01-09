package goql

// Context A Context structure for holding the state of the entire server
type Context struct {
	dbList map[string]*DB
}

// DB Database structure which holds all the tables.
type DB struct {
}
