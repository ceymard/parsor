package parsor

import (
	"fmt"
	"strings"
)

// Text returns the concatenated text of all the positions inside this slice
func (ps *PositionSlice) Text() string {
	var res = ""
	for _, s := range *ps {
		res += s.Text
	}
	return res
}

// OriginalText returns the concatenated, unmodified text of all the positions inside
// this slice.
func (ps *PositionSlice) OriginalText() string {
	var res = ""
	for _, s := range *ps {
		res += s.OriginalText
	}
	return res
}

func (ps *PositionSlice) Line() int {
	return (*ps)[0].Line
}

func (ps *PositionSlice) Column() int {
	return (*ps)[0].Column
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

// Until returns all the positions until reaching a certain token.
// It includes skipped tokens in its result.
func (p *Position) Until(tokens ...*Token) (*PositionSlice, error) {
	var res = make([]*Position, 10)

	return nil, nil
}

// While returns all the positions matching the asked tokens
func (p *Position) While(tokens ...*Token) (*PositionSlice, error) {
	if len(tokens) == 0 {
		return nil, nil
	}

	return nil, nil
}

func (p *Position) getNext(skip bool, tokens ...*Token) (*Position, error) {
	var n, err = p.lexer.getNextPosition(p, skip)

	if err != nil {
		return nil, err
	}

	if n == nil {
		if len(tokens) > 0 {
			return nil, fmt.Errorf("reached end of input")
		}
		return nil, nil
	}

	// return an error if the next token doesn't match
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

	return n, nil
}

// Next gives the next position that optionally matches the provided tokens.
// It skips "skippables" by default.
func (p *Position) Next(tokens ...*Token) (*Position, error) {
	return p.getNext(true, tokens...)
}

// NextNoSkip Does the same as Next without skipping the skipables
func (p *Position) NextNoSkip(tokens ...*Token) (*Position, error) {
	return p.getNext(false, tokens...)
}

func (p *Position) Unexpected() error {
	return fmt.Errorf("unexpected %s at line %d column %d", p.Token.Name, p.Line, p.Column)
}
