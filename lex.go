package parsor

import (
	"fmt"
	"io"
	"strings"
)

func (lex *Lexer) getNextPosition(current *Position, skip bool) (*Position, error) {
	var more = make([]byte, 1024)
	var fillUpMore = false
	var err error
	var read int

	for {
		if lex.atEOF && len(lex.buffer) == 0 {
			return nil, nil
		}

		if len(lex.buffer) == 0 || fillUpMore {
			// we should check for EOF
			read, err = lex.reader.Read(more)

			if err == io.EOF {
				lex.atEOF = true
				if len(lex.buffer) == 0 {
					// no next position
					return nil, nil
				}
			}

			if err != nil && err != io.EOF {
				return nil, err
			}

			lex.buffer += string(more[:read])
			// fmt.Println(lex.buffer)
			fillUpMore = false
		}

		for _, tok := range lex.tokens {
			var matches = tok.Regexp.FindStringSubmatch(lex.buffer)

			// No match found !
			if matches == nil {
				continue
			}

			if len(matches[0]) == len(lex.buffer) && !lex.atEOF {
				// the match goes to the end of our buffer.
				// for all we know, the token spans more than what we just have
				// in our buffer, so we'll ask to read more
				fillUpMore = true
				break
			}

			// FIXME, this is where we update Line and Column
			// We found a token ! yaay !
			var str = matches[0]
			if tok.Transformer != nil {
				str = (*tok.Transformer)(matches)
			}
			lex.buffer = lex.buffer[len(str):]
			var position = &Position{
				Text:         str,
				OriginalText: matches[0],
				Token:        tok,
				Line:         lex.line,
				Column:       lex.column,
				lexer:        lex,
			}

			var newlines = strings.Count(str, "\n")
			if newlines > 0 {
				lex.column = len(str) - strings.LastIndex(str, "\n")
			} else {
				lex.column += len(str)
			}
			lex.line += newlines

			if current != nil {
				current.next = position
			}
			current = position

			// If we need skipping then we continue matching the input
			if skip && current.Token.IsSkippable {
				continue
			}

			return current, nil
		}

		if !fillUpMore {
			// if we got here, then it means that we didn't match any token
			// and yet we're not trying to fetch more input.
			// this is an error !
			var l = len(lex.buffer)
			if l > 16 {
				l = 16
			}
			return nil, fmt.Errorf("unexpected input at %d %d `%s`...", lex.line, lex.column, lex.buffer[:l])
		}
	}

}
