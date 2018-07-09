package parsor

import (
	"fmt"
	"strings"
)

func (ps *PositionSlice) Text() string {
	var res = ""
	for _, s := range *ps {
		res += s.Text
	}
	return res
}

func (ps *PositionSlice) OriginalText() string {
	var res = ""
	for _, s := range *ps {
		res += s.OriginalText
	}
	return res
}

func (ps *PositionSlice) Last() *Position {
	return (*ps)[len(*ps)-1]
}

func (ps *PositionSlice) Next(toks ...*Token) (*Position, error) {
	return ps.Last().Next(toks...)
}

func (ps *PositionSlice) Skip(toks ...*Token) (*Position, error) {
	return ps.Last().Skip(toks...)
}

// Is tells if the position corresponds to a given token
func (p *Position) Is(toks ...*Token) bool {
	// if no token provided, then this position is whatever we want.
	if len(toks) == 0 {
		return true
	}

	for _, t := range toks {
		if t == p.Token {
			return true
		}
	}
	return false
}

func (p *Position) Skip(tokens ...*Token) (*Position, error) {
	return nil, nil
}

func (p *Position) Until(tokens ...*Token) (*PositionSlice, error) {
	return nil, nil
}

func (p *Position) While(tokens ...*Token) (*PositionSlice, error) {
	return nil, nil
}

// Get the next position from the current one
func (p *Position) _next() (*Position, error) {
	if p.next != nil {
		return p.next, nil
	}

	// no next, so we ask the Lexer
	if newnext, err := p.lexer.getNextPosition(); err != nil {
		return nil, err
	} else {
		p.next = newnext
		return newnext, nil
	}
}

func (p *Position) _nextSkipped() (*Position, error) {
	var n, err = p._next()
	if err != nil {
		return nil, err
	}
	var sk = p.lexer.skips
	for n != nil && n.Is(sk...) {
		var n2, err2 = n._next()
		if err2 != nil {
			return nil, err2
		}
		n = n2
	}
	return n, nil
}

// Next
func (p *Position) Next(tokens ...*Token) (*Position, error) {
	var n, err = p._nextSkipped()
	if err != nil {
		return nil, err
	}

	if !n.Is(tokens...) {
		var names = make([]string, 0)
		for _, t := range tokens {
			names = append(names, t.Name)
		}
		return nil, fmt.Errorf(`at line %d and position %d got %s but was expecting %s`,
			n.Line,
			n.Column,
			n.Token.Name,
			strings.Join(names, " or "),
		)
	}

	return nil, nil
}
