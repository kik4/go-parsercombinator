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
		parsers []Parser
	}

	// RuleFunc is function validates parse rule.
	RuleFunc func(string) (string, int, bool)

	// CountFunc is function counts RuleFunc.
	CountFunc func(string) (string, int, error)
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
func (t *Terminal) Parse(s string) (string, int, error) {
	return t.f(s)
}

// Parse executes parser
func (nt *NonTerminal) Parse(s string) (string, int, error) {
	content := make([]byte, 0)
	read := 0

	for _, p := range nt.parsers {
		str, num, err := p.Parse(s[read:])
		if err != nil {
			return "", 0, err
		}
		read += num
		content = append(content, str...)
	}
	return string(content), read, nil
}

// Sequence combines Parsers
func Sequence(args ...Parser) *NonTerminal {
	return &NonTerminal{args}
}
