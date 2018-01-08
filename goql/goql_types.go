package goql

// SQLCommand SQL command types
type SQLCommand int32

const (
	// SQLCommandInsert Insert command from SQL
	SQLCommandInsert SQLCommand = iota
	// SQLCommandSelect Select command from SQL
	SQLCommandSelect SQLCommand = iota
)
