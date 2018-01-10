package goql

// Statement default statement interface
type Statement interface {
	Type() SQLCommand
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

func (c CreateDatabaseStmt) String() string {
	return "CREATE DATABASE " + c.DbName
}

// Type type of the command
func (c CreateDatabaseStmt) Type() SQLCommand {
	return SQLCommandCreateDatabase
}

// Column structure representing the column in a table
type Column struct {
	Name string
	Type SQLColumnDataType
	Size int32
}

// CreateTableStmt Create Table Statement
type CreateTableStmt struct {
	exists    bool
	TableName string
	Columns   *[]Column
}

func (c CreateTableStmt) String() string {
	return "CREATE TABLE "
}

// Type type of the command
func (c CreateTableStmt) Type() SQLCommand {
	return SQLCommandCreateTable
}
