package goqlparser_test

import (
	"strings"
	"testing"

	goqlparser "github.com/ssarangi/goql/goql/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok goqlparser.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: goqlparser.EOF},
		{s: `#`, tok: goqlparser.ILLEGAL, lit: `#`},
		{s: ` `, tok: goqlparser.WS, lit: " "},
		{s: "\t", tok: goqlparser.WS, lit: "\t"},
		{s: "\n", tok: goqlparser.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: goqlparser.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: goqlparser.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: goqlparser.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: goqlparser.FROM, lit: "FROM"},
		{s: `SELECT`, tok: goqlparser.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := goqlparser.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
