package parsor

import (
	"io"
	"strings"
)

func NewParser(input io.Reader, tokens []*Token, autoSkipTokens []*Token) (*Position, error) {
	// first build a lexer
	var lex = Lexer{tokens, autoSkipTokens, input, ""}

	// then from the lexer, get the first position

	return nil, nil
}

func NewParserString(input string, tokens []*Token, autoSkipTokens []*Token) (*Position, error) {
	return NewParser(strings.NewReader(input), tokens, autoSkipTokens)
}
