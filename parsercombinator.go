package parsercombinator

import (
	"errors"
)

type (
	// Parser is parse executer.
	Parser interface {
		Parse(s string) (string, int, error)
	}

	// Terminal is terminal symbol parser.
	Terminal struct {
		f CountFunc
	}

	// NonTerminal is nonterminal symbol parser.
	NonTerminal struct {
		f CombineFunc
	}

	// RuleFunc is function validates parse rule.
	RuleFunc func(string) (string, int, bool)

	// CountFunc is function counts RuleFunc.
	CountFunc func(string) (string, int, error)

	// CombineFunc is function conbines parser.
	CombineFunc func(string) (string, int, error)
)

// Once adapts parse rule once.
func (rf RuleFunc) Once() Parser {
	return &Terminal{f: func(test string) (string, int, error) {
		str, num, succeeded := rf(test)
		if succeeded {
			return str, num, nil
		}
		return "", 0, errors.New("once is failed")
	}}
}

// Parse executes parser
func (t *Terminal) Parse(test string) (string, int, error) {
	return t.f(test)
}

// Parse executes parser
func (nt *NonTerminal) Parse(test string) (string, int, error) {
	return nt.f(test)
}

// Sequence combines Parsers
func Sequence(parsers ...Parser) *NonTerminal {
	return &NonTerminal{func(test string) (string, int, error) {
		content := make([]byte, 0)
		read := 0

		for _, p := range parsers {
			str, num, err := p.Parse(test[read:])
			if err != nil {
				return "", 0, err
			}
			read += num
			content = append(content, str...)
		}
		return string(content), read, nil
	}}
}

// Or selects matched parse result
func Or(p1, p2 Parser) *NonTerminal {
	return &NonTerminal{func(test string) (string, int, error) {
		str, num, err := p1.Parse(test)
		if err == nil {
			return str, num, err
		}

		str, num, err = p2.Parse(test)
		if err == nil {
			return str, num, err
		}

		return "", 0, errors.New("NonTerminal parse is failed")
	}}
}
