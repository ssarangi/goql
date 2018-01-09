package goql

import "log"

// Context A Context structure for holding the state of the entire server
type Context struct {
	dbList map[string]*DB
}

// Execute execute the current statement
func (c *Context) Execute(stmt Statement) {
	switch t := stmt.(type) {
	case CreateDatabaseStmt:
		c.executeCreateDatabase(t)
	}
}

func (c *Context) executeCreateDatabase(stmt CreateDatabaseStmt) {
	log.Println("Executed Command: Create Database " + stmt.DbName)
	c.dbList[stmt.DbName] = &DB{name: stmt.DbName}
}

// DB Database structure which holds all the tables.
type DB struct {
	name string
}
