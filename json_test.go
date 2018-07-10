package parsor

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	tLBracket = NewToken("{", `\{`)
	tRBracket = NewToken("}", `\}`)
	tLBrace   = NewToken("[", `\[`)
	tRBrace   = NewToken("]", `\]`)
	tNumber   = NewToken("number", `\d+(\.\d+)?`)
	tString   = NewTokenAction("string", `"(\\"|[^"])*"`, func(match []string) string {
		var str = match[0]
		return strings.Replace(str[1:len(str)-1], `\`+string(str[0]), string(str[0]), -1)
	})
	tTrue  = NewToken("true", `true`)
	tFalse = NewToken("false", `false`)
	tNull  = NewToken("null", `null`)
	tWs    = NewSkippableToken("whitespace", `[ \t\n\r]+`)
)

func parseValue(p *Position) (interface{}, *Position, error) {

	if p, err := p.Next(tTrue); err == nil {
		return true, p, nil
	} else if p, err := p.Next(tFalse); err == nil {
		return false, p, nil
	} else if p, err := p.Next(tNull); err == nil {
		return nil, p, nil
	} else if p, err := p.Next(tLBrace); err == nil {
		return parseArray(p)
	} else if p, err := p.Next(tLBracket); err == nil {
		return parseObject(p)
	} else if p, err := p.Next(tString); err == nil {
		return p.Text, p, nil
	} else if p.Is(tNumber) {
		var f, err = strconv.ParseFloat(p.Text, 64)
		return f, p, err
	} else {
		return nil, nil, p.Unexpected()
	}
}

func parseArray(p *Position) ([]interface{}, *Position, error) {
	if !p.Is(tLBracket) {
		return nil, nil, fmt.Errorf("arrays must start with [")
	}

	for {
		if p.Is(tRBracket) {
			// we're done when we reach the RBracket
			break
		}
	}

	return nil, nil, nil
}

func parseObject(p *Position) (map[string]interface{}, *Position, error) {
	return nil, nil, nil
}

func parse(str string) (interface{}, error) {
	var pos, err = NewParserString(str, []*Token{
		tLBracket,
		tRBracket,
		tLBrace,
		tRBrace,
		tNumber,
		tString,
		tTrue,
		tFalse,
		tNull,
		tWs,
	})
	if err != nil {
		return nil, err
	}
	return parseValue(pos)
}
