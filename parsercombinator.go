package parsercombinator

import (
	"errors"
)

type (
	// Parser is parse executer.
	Parser struct {
		f ParseFunc
	}

	// RuleFunc is function validates parse rule.
	RuleFunc func(string) (string, int, bool)

	// ParseFunc is function parses.
	ParseFunc func(string) (string, int, error)
)

// Parse executes parser
func (p *Parser) Parse(test string) (string, int, error) {
	return p.f(test)
}

// Sequence combines Parsers
func Sequence(parsers ...*Parser) *Parser {
	return &Parser{func(test string) (string, int, error) {
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
func Or(p1, p2 *Parser) *Parser {
	return &Parser{func(test string) (string, int, error) {
		str, num, err := p1.Parse(test)
		if err == nil {
			return str, num, err
		}

		str, num, err = p2.Parse(test)
		if err == nil {
			return str, num, err
		}

		return "", 0, errors.New("Or parse is failed")
	}}
}
