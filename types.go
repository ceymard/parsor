package parsor

import (
	"io"
	"regexp"
)

// Token is tried in the input.
type Token struct {
	// Name of the Token for debug and errors
	Name string

	// Regexp is used to match text in the input
	Regexp regexp.Regexp

	// Transformer if set transforms the matched text into
	// something else, for instance if you match a string but
	// want the contents of the Lexeme to only contain the insides
	Transformer *func([]string) string

	IsSkippable bool
}

// Position is a piece of matched text that corresponds to a Token
type Position struct {
	// Text contains what was matched post-modification by Token.Transformer
	Text string

	// OriginalText contains all the Lexeme's text before modification
	// by Token's Transformer if it had one.
	OriginalText string

	// Line at which the lexeme was found
	Line int

	// Column at which the lexeme begins
	Column int

	// The Token that generated this match
	Token *Token

	lexer *Lexer
	next  *Position
}

// PositionSlice is a slice of positions
type PositionSlice []*Position

// Lexer reads from a Reader and outputs Lexemes.
type Lexer struct {
	tokens []*Token
	reader io.Reader
	buffer string
	atEOF  bool
	line   int
	column int
}
