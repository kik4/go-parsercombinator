package parsercombinator

import (
	"errors"
	"unicode"
	"unicode/utf8"
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

// AnyChar reads a charater.
func AnyChar() RuleFunc {
	return func(test string) (string, int, bool) {
		if utf8.RuneCountInString(test) < 1 {
			return "", 0, false
		}

		return test[:1], 1, true
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

// Digit read a digit.
func Digit() RuleFunc {
	return func(test string) (string, int, bool) {
		if utf8.RuneCountInString(test) < 1 {
			return "", 0, false
		}

		c := test[:1][0]
		if unicode.IsDigit(rune(c)) {
			return test[:1], 1, true
		}
		return "", 0, false
	}
}

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
