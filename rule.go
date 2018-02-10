package parsercombinator

import (
	"unicode"
	"unicode/utf8"
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

// Letter read a Letter.
func Letter() RuleFunc {
	return func(test string) (string, int, bool) {
		if utf8.RuneCountInString(test) < 1 {
			return "", 0, false
		}

		c := test[:1][0]
		if unicode.IsLetter(rune(c)) {
			return test[:1], 1, true
		}
		return "", 0, false
	}
}
