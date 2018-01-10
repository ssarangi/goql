package goqlparser

// https://blog.gopheracademy.com/advent-2014/parsers-lexers/

import (
	"fmt"
	"log"
	"strings"

	"github.com/ssarangi/goql/goql"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct {
	Fields    []string
	TableName string
}

// Parser represents a parser
type Parser struct {
	command string
	s       *Scanner
	buf     struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(command string) *Parser {
	return &Parser{s: NewScanner(strings.NewReader(command))}
}

// Parse parses a SQL statement.
func (p *Parser) Parse() (goql.Statement, error) {
	var stmt goql.Statement
	var err error
	if tok, _ := p.scanIgnoreWhitespace(); tok == CREATE {
		if tok, _ := p.scanIgnoreWhitespace(); tok == DATABASE {
			stmt, err = p.parseCreateDatabase()
		} else if tok == TABLE {
			stmt, err = p.parseCreateTable()
		}
	}

	return stmt, err
}

func (p *Parser) parseCreateDatabase() (*goql.CreateDatabaseStmt, error) {
	// Now parse the database name
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("Invalid Identifier passed to 'Create Database " + lit + "'")
	}

	stmt := new(goql.CreateDatabaseStmt)
	stmt.DbName = lit
	log.Println("Created Database: " + lit)
	return stmt, nil
}

func (p *Parser) parseColumnName() (string, error) {
	// Parse the column name
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return "", fmt.Errorf("Expected COLUMN NAME after CREATE TABLE <table name>")
	}

	return lit, nil
}

func (p *Parser) parseColumnDataType() (goql.SQLColumnDataType, uint32, error) {
	tok, _ := p.scanIgnoreWhitespace()
	switch tok {
	case VARCHAR:
		// Now read the size
		tok, _ = p.scanIgnoreWhitespace()
		if tok != LEFT_BRACKET {
			return goql.SQLColumnUNKNOWN, 0, fmt.Errorf("No size specified for VARCHAR: %s", p.command)
		}
		tok, _ = p.scanIgnoreWhitespace()

		if tok != RIGHT_BRACKET {
			return goql.SQLColumnUNKNOWN, 0, fmt.Errorf("Incorrect size format: %s", p.command)
		}
		break
	case INT:
		break
	}

	return goql.SQLColumnUNKNOWN, 0, fmt.Errorf("Invalid Column datatype specified: %s Command: %s", tok, p.command)
}

func (p *Parser) parseSingleColumnDefinition() (*goql.TableColumn, error) {
	columnName, err := p.parseColumnName()
	if err != nil {
		return nil, err
	}

	columnDataType, size, err := p.parseColumnDataType()
	c := &goql.TableColumn{Name: columnName, Type: columnDataType, Size: size}

	return c, nil
}

func (p *Parser) parseColumnDefinitions() ([]*goql.TableColumn, error) {
	// Find the column definition
	tok, _ := p.scanIgnoreWhitespace()
	if tok != LEFT_BRACKET {
		return nil, fmt.Errorf("Expected ( after TABLE NAME")
	}

	var columns []*goql.TableColumn

	for tok != RIGHT_BRACKET {
		columnDef, err := p.parseSingleColumnDefinition()
		if err != nil {
			return nil, err
		}
		columns = append(columns, columnDef)
	}

	return columns, nil
}

func (p *Parser) parseCreateTable() (*goql.CreateTableStmt, error) {
	// We have already parsed till Table. Now parse the table name
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("Expected TABLE_NAME after CREATE TABLE command but got %s: %s", lit, p.command)
	}

	p.parseColumnDefinitions()
	return nil, nil
}

// Parse parses a SQL SELECT statement.
func (p *Parser) parseStatement() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	// First token should be a 'SELECT' keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", lit)
	}

	// Next we should loop over all our comma-delimited fields.
	for {
		// Read a field.
		tok, lit := p.scanIgnoreWhitespace()
		if tok != IDENT && tok != ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}

		stmt.Fields = append(stmt.Fields, lit)

		// If the next token is not a comma then break the loop.
		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}

	// Next we should see the 'FROM' keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	// Finally we should read the table name
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected table name", lit)
	}
	stmt.TableName = lit

	// Return the successfully parsed statement.
	return stmt, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// Peek to the next character
func (p *Parser) peekIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	scanCount := 0

	for tok == WS {
		tok, lit = p.scan()
		scanCount++
	}

	for i := 0; i < scanCount; i++ {
		p.unscan()
	}

	return
}
