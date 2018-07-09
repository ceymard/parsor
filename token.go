package parsor

import (
	"regexp"
)

// NewToken returns a new token ready for inclusion in the parser
func NewToken(name string, expression string, action *func([]string) string) *Token {
	return &Token{name, *regexp.MustCompile("^" + expression), action}
}
