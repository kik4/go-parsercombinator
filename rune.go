package parsercombinator

import (
	"unicode"
)

type runeValidationFunc func(rune) bool

// AnyRune reads a rune.
func AnyRune() RuneFunc {
	fn := func(r rune) bool {
		return true
	}
	return createCommonRuneFunc(fn)
}

// Digit read a digit.
func Digit() RuneFunc {
	return createCommonRuneFunc(unicode.IsDigit)
}

// Letter read a Letter.
func Letter() RuneFunc {
	return createCommonRuneFunc(unicode.IsLetter)
}

// Rune read a rune assigned
func Rune(needle rune) RuneFunc {
	fn := func(r rune) bool {
		return r == needle
	}
	return createCommonRuneFunc(fn)
}

// Control rune read
func Control() RuneFunc {
	return createCommonRuneFunc(unicode.IsControl)
}

// Graphic rune read
func Graphic() RuneFunc {
	return createCommonRuneFunc(unicode.IsGraphic)
}

// Lower rune read
func Lower() RuneFunc {
	return createCommonRuneFunc(unicode.IsLower)
}

// Mark rune read
func Mark() RuneFunc {
	return createCommonRuneFunc(unicode.IsMark)
}

// Number rune read
func Number() RuneFunc {
	return createCommonRuneFunc(unicode.IsNumber)
}

// Print rune read
func Print() RuneFunc {
	return createCommonRuneFunc(unicode.IsPrint)
}

// Punct rune read
func Punct() RuneFunc {
	return createCommonRuneFunc(unicode.IsPunct)
}

// Space rune read '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func Space() RuneFunc {
	return createCommonRuneFunc(unicode.IsSpace)
}

// Symbol rune read
func Symbol() RuneFunc {
	return createCommonRuneFunc(unicode.IsSymbol)
}

// Title rune read
func Title() RuneFunc {
	return createCommonRuneFunc(unicode.IsTitle)
}

// Upper rune read
func Upper() RuneFunc {
	return createCommonRuneFunc(unicode.IsUpper)
}

// In rune read
func In(ranges ...*unicode.RangeTable) RuneFunc {
	fn := func(r rune) bool {
		return unicode.In(r, ranges...)
	}
	return createCommonRuneFunc(fn)
}

// InStr rune read in assgined string
func InStr(table string) RuneFunc {
	fn := func(r rune) bool {
		for _, test := range table {
			if r == test {
				return true
			}
		}
		return false
	}
	return createCommonRuneFunc(fn)
}

func createCommonRuneFunc(fn runeValidationFunc) RuneFunc {
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
