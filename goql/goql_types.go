package goql

// MetaCommandResult Result type for whether a meta command was successfully executed or not
type MetaCommandResult int32

const (
	// MetaCommandSuccess Successfully parsed meta command state
	MetaCommandSuccess MetaCommandResult = iota
	// MetaCommandUnrecognized Invalid meta command state passed.
	MetaCommandUnrecognized
)

// SQLCommand SQL command types
type SQLCommand int32

const (
	// SQLCommandInsert Insert command from SQL
	SQLCommandInsert SQLCommand = iota
	// SQLCommandSelect Select command from SQL
	SQLCommandSelect
)

// PrepareSQLCommandResult Prepare SQL command into structure
type PrepareSQLCommandResult int32

const (
	// PrepareSuccess Successfully parsed the SQL command.
	PrepareSuccess PrepareSQLCommandResult = iota
	// PrepareUnrecognizedStatement Didn't recognize the statement.
	PrepareUnrecognizedStatement
)

// SQLStatement A statement containing the SQL command
type SQLStatement struct {
	// CommandType The type of the SQL Command
	CommandType SQLCommand
}
