package parsor

import (
	"io"
	"strings"
)

func NewParser(input io.Reader, tokens []*Token) (*Position, error) {
	// first build a lexer
	var lex = Lexer{
		tokens: tokens,
		reader: input,
	}

	return lex.getNextPosition(nil, true)
}

func NewParserString(input string, tokens []*Token) (*Position, error) {
	return NewParser(strings.NewReader(input), tokens)
}
