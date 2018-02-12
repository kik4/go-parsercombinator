package parsercombinator

import (
	"unicode"
)

type simpleIsFunc func(rune) bool

// AnyChar reads a charater.
func AnyChar() RuleFunc {
	fn := func(r rune) bool {
		return true
	}
	return createCommonRuleFunc(fn)
}

// String validates equal test stiring.
func String(needle string) RuleFunc {
	return func(test string) (string, int, bool) {
		length := len(needle)
		if len(test) < length {
			return "", 0, false
		}
		if test[:length] == needle {
			return test[:length], length, true
		}
		return "", 0, false
	}
}

// Digit read a digit.
func Digit() RuleFunc {
	return createCommonRuleFunc(unicode.IsDigit)
}

// Letter read a Letter.
func Letter() RuleFunc {
	return createCommonRuleFunc(unicode.IsLetter)
}

// Char read a rune assigned
func Char(needle rune) RuleFunc {
	fn := func(r rune) bool {
		return r == needle
	}
	return createCommonRuleFunc(fn)
}

func createCommonRuleFunc(fn simpleIsFunc) RuleFunc {
	return func(test string) (string, int, bool) {
		for _, r := range test {
			if !fn(r) {
				return "", 0, false
			}
			str := string(r)
			return str, len(str), true
		}
		return "", 0, false
	}
}
