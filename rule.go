package parsercombinator

import (
	"unicode"
)

type simpleIsFunc func(rune) bool

// AnyRune reads a rune.
func AnyRune() RuleFunc {
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

// Rune read a rune assigned
func Rune(needle rune) RuleFunc {
	fn := func(r rune) bool {
		return r == needle
	}
	return createCommonRuleFunc(fn)
}

// Control rune read
func Control() RuleFunc {
	return createCommonRuleFunc(unicode.IsControl)
}

// Graphic rune read
func Graphic() RuleFunc {
	return createCommonRuleFunc(unicode.IsGraphic)
}

// Lower rune read
func Lower() RuleFunc {
	return createCommonRuleFunc(unicode.IsLower)
}

// Mark rune read
func Mark() RuleFunc {
	return createCommonRuleFunc(unicode.IsMark)
}

// Number rune read
func Number() RuleFunc {
	return createCommonRuleFunc(unicode.IsNumber)
}

// Print rune read
func Print() RuleFunc {
	return createCommonRuleFunc(unicode.IsPrint)
}

// Punct rune read
func Punct() RuleFunc {
	return createCommonRuleFunc(unicode.IsPunct)
}

// Space rune read '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func Space() RuleFunc {
	return createCommonRuleFunc(unicode.IsSpace)
}

// Symbol rune read
func Symbol() RuleFunc {
	return createCommonRuleFunc(unicode.IsSymbol)
}

// Title rune read
func Title() RuleFunc {
	return createCommonRuleFunc(unicode.IsTitle)
}

// Upper rune read
func Upper() RuleFunc {
	return createCommonRuleFunc(unicode.IsUpper)
}

// In rune read
func In(ranges ...*unicode.RangeTable) RuleFunc {
	fn := func(r rune) bool {
		return unicode.In(r, ranges...)
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
