package parsor

import (
	"regexp"
)

// NewSkippableToken a token that may be skipped by the lexer
func NewSkippableToken(name string, expression string) *Token {
	return &Token{name, *regexp.MustCompile("^" + expression), nil, true}
}

// NewToken returns a new token ready for inclusion in the parser
func NewToken(name string, expression string) *Token {
	return &Token{name, *regexp.MustCompile("^" + expression), nil, false}
}

// NewTokenAction returns a token that has an action
func NewTokenAction(name string, expression string, action func([]string) string) *Token {
	return &Token{name, *regexp.MustCompile("^" + expression), &action, false}
}
