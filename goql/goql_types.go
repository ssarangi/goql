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
	// SQLCommandCreateTable SQL Create Table Command
	SQLCommandCreateTable SQLCommand = iota
	// SQLCommandCreateDatabase SQL Create Database Command
	SQLCommandCreateDatabase
	// SQLCommandInsert Insert command from SQL
	SQLCommandInsert
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

// SQLConstraint Value type representing a SQL constraint.
type SQLConstraint int32

// SQLConstraints List of SQL constraints
const (
	// SQLConstraintNOTNULL Not Null constraint.
	SQLConstraintNOTNULL SQLConstraint = iota
	// SQLConstraintUNIQUE Unique constraint.
	SQLConstraintUNIQUE
	// SQLConstraintPRIMARYKEY Primary Key constraint
	SQLConstraintPRIMARYKEY
	// SQLConstraintFOREIGNKEY Foreign Key constraint
	SQLConstraintFOREIGNKEY
	// SQLConstraintCHECK Check constraint
	SQLConstraintCHECK
	// SQLConstraintDEFAULT Default constraint
	SQLConstraintDEFAULT
	// SQLConstraintINDEX Index constraint
	SQLConstraintINDEX
)
