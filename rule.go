package parsercombinator

import (
	"unicode"
)

// AnyChar reads a charater.
func AnyChar() RuleFunc {
	return func(test string) (string, int, bool) {
		for _, r := range test {
			str := string(r)
			return str, len(str), true
		}
		return "", 0, false
	}
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
	return func(test string) (string, int, bool) {
		for _, r := range test {
			if !unicode.IsDigit(r) {
				return "", 0, false
			}
			str := string(r)
			return str, len(str), true
		}
		return "", 0, false
	}
}

// Letter read a Letter.
func Letter() RuleFunc {
	return func(test string) (string, int, bool) {
		for _, r := range test {
			if !unicode.IsLetter(r) {
				return "", 0, false
			}
			str := string(r)
			return str, len(str), true
		}
		return "", 0, false
	}
}

// Char read a rune assigned
func Char(needle rune) RuleFunc {
	return func(test string) (string, int, bool) {
		for _, r := range test {
			if r != needle {
				return "", 0, false
			}
			str := string(r)
			return str, len(str), true
		}
		return "", 0, false
	}
}
