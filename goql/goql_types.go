package goql

// MetaCommandResult represents the type for enums which are used for handling commands like ".exit"
type MetaCommandResult int

const (
	// MetaCommandSuccess Successfully parsed meta command state
	MetaCommandSuccess = 1 << iota
	// MetaCommandUnrecognized Invalid meta command state passed.
	MetaCommandUnrecognized
)
