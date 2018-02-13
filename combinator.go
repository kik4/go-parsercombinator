package parsercombinator

import (
	"errors"
)

type (
	// Parser is parse executer.
	Parser struct {
		f ParseFunc
	}

	// RuneFunc is function validates parse rule.
	RuneFunc func([]rune) (rune, int, bool)

	// ParseFunc is function parses.
	ParseFunc func([]rune) (interface{}, int, error)

	// SelectFunc is function selects new value.
	SelectFunc func([]interface{}) interface{}
)

// Parse executes parser
func (p *Parser) Parse(target string) (interface{}, int, error) {
	return p.f([]rune(target))
}

// Sequence combines Parsers
func Sequence(parsers []*Parser, selector SelectFunc) *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		params := []interface{}{}
		read := 0

		for _, p := range parsers {
			param, num, err := p.f(test[read:])
			read += num
			if err != nil {
				return nil, read, err
			}
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
	return &Parser{func(test []rune) (interface{}, int, error) {
		val, num, err := p1.f(test)
		if err == nil {
			return val, num, err
		}

		val, num, err = p2.f(test)
		if err == nil {
			return val, num, err
		}

		return nil, num, errors.New("Or parse is failed")
	}}
}
