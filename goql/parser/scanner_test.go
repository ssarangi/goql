package goql_test

import (
	"strings"
	"testing"

	goql "github.com/ssarangi/goql/goql/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok goql.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: goql.EOF},
		{s: `#`, tok: goql.ILLEGAL, lit: `#`},
		{s: ` `, tok: goql.WS, lit: " "},
		{s: "\t", tok: goql.WS, lit: "\t"},
		{s: "\n", tok: goql.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: goql.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: goql.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: goql.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: goql.FROM, lit: "FROM"},
		{s: `SELECT`, tok: goql.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := goql.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
