package goql

const (
	// MetaCommandSuccess Successfully parsed meta command state
	MetaCommandSuccess = 1 << iota
	// MetaCommandUnrecognized Invalid meta command state passed.
	MetaCommandUnrecognized
)

const (
	// SQLCommandInsert Insert command from SQL
	SQLCommandInsert = 1 << iota
	// SQLCommandSelect Select command from SQL
	SQLCommandSelect
)
