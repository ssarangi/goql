package goqlparser

// Token represents a lexical token
type Token int

const (
	// Special Tokens
	// ILLEGAL token
	ILLEGAL Token = iota
	// EOF End of File
	EOF
	// WS Whitespace
	WS

	// Literals
	// IDENT Identifier
	IDENT // fields, table_name

	// Misc characters
	// ASTERISK * character
	ASTERISK // *
	// COMMA , character
	COMMA // ,
	// LEFT_BRACKET left bracket
	LEFT_BRACKET // (
	// RIGHT_BRACKET right bracket
	RIGHT_BRACKET // )
	// SEMICOLON semi colon
	SEMICOLON

	// Keywords
	// SELECT Keyword
	SELECT
	// FROM Keyword
	FROM
	// CREATE Keyword
	CREATE
	// DATABASE Keyword
	DATABASE
	// TABLE Keyword
	TABLE
	// INSERT Keyword
	INSERT

	// DATATypes
	// VARCHAR variable characters
	VARCHAR
	// INT integer
	INT
	// BOOLEAN bool
	BOOLEAN
)
