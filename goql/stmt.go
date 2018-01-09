package goql

// Statement default statement interface
type Statement interface {
	String() string
}

// type execCtx struct {
// 	mu sync.RWMutex
// }

// type col struct {
// 	name       string
// 	constraint SQLConstraint
// }

// // createTableStmt represents a SQL Create Table statement.
// type createTableStmt struct {
// 	exists    bool
// 	cols      []*col
// 	tableName string
// }

// CreateDatabaseStmt Create Database Statement
type CreateDatabaseStmt struct {
	exists bool
	DbName string
}

func (c *CreateDatabaseStmt) String() string {
	return "CREATE DATABASE " + c.DbName
}
