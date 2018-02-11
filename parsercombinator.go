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
	ParseFunc func(string) (interface{}, int, error)

	// SelectFunc is function selects new value.
	SelectFunc func([]interface{}) interface{}
)

// Parse executes parser
func (p *Parser) Parse(test string) (interface{}, int, error) {
	return p.f(test)
}

// Sequence combines Parsers
func Sequence(selector SelectFunc, parsers ...*Parser) *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		params := []interface{}{}
		read := 0

		for _, p := range parsers {
			param, num, err := p.Parse(test[read:])
			if err != nil {
				return nil, 0, err
			}
			read += num
			params = append(params, param)
		}

		if selector == nil {
			return params, read, nil
		}

		return selector(params), read, nil
	}}
}

// Or selects matched parse result
func Or(p1, p2 *Parser) *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		val, num, err := p1.Parse(test)
		if err == nil {
			return val, num, err
		}

		val, num, err = p2.Parse(test)
		if err == nil {
			return val, num, err
		}

		return nil, 0, errors.New("Or parse is failed")
	}}
}
