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

	// Keywords
	// SELECT Keyword
	SELECT
	// FROM Keyword
	FROM
)
