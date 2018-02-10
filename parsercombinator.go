package parsercombinator

import (
	"errors"
	"unicode/utf8"
)

type (
	// Parser is parse executer.
	Parser struct {
		f CountFunc
	}

	// RuleFunc is function validates parse rule.
	RuleFunc func(string) (string, int, bool)

	// CountFunc is function counts RuleFunc.
	CountFunc func(string) (string, error)
)

// AnyChar reads a charater.
func AnyChar() RuleFunc {
	return func(test string) (string, int, bool) {
		if utf8.RuneCountInString(test) >= 1 {
			return test[:1], 1, true
		}
		return "", 0, false
	}
}

// String validates equal test stiring.
func String(s string) RuleFunc {
	return func(test string) (string, int, bool) {
		num := utf8.RuneCountInString(s)
		if utf8.RuneCountInString(test) < num {
			return "", 0, false
		}
		if test[0:num] == s {
			return test[0:num], num, true
		}
		return "", 0, false
	}
}

// Once adapts parse rule once.
func (rf RuleFunc) Once() *Parser {
	return &Parser{f: func(test string) (string, error) {
		str, _, succeeded := rf(test)
		if succeeded {
			return str, nil
		}
		return "", errors.New("hoge")
	}}
}

// Parse executes parser
func (p *Parser) Parse(s string) (string, error) {
	return p.f(s)
}
