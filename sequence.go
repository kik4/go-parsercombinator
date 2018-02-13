package parsercombinator

import (
	"errors"
)

// Once adapts parse rule once.
func (rf RuneFunc) Once() *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		r, num, ok := rf(test)
		if ok {
			return string(r), num, nil
		}
		return "", num, errors.New("Once is failed")
	}}
}

// AtLeastOnce adapts parse rule once or more.
func (rf RuneFunc) AtLeastOnce() *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		content := []rune{}
		read := 0
		r, num, ok := rf(test[read:])
		for ; ok; r, num, ok = rf(test[read:]) {
			read += num
			content = append(content, r)
		}

		if read == 0 {
			return "", num, errors.New("AtLeastOnce is failed")
		}

		return string(content), read, nil
	}}
}

// Many adapts parse rule 0 or more times.
func (rf RuneFunc) Many() *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		content := []rune{}
		read := 0
		r, num, ok := rf(test[read:])
		for ; ok; r, num, ok = rf(test[read:]) {
			read += num
			content = append(content, r)
		}

		return string(content), read, nil
	}}
}

// Repeat adapts parse rule count times.
func (rf RuneFunc) Repeat(count int) *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		if count <= 0 {
			return "", 0, errors.New("Repeat needs 1 or more times")
		}

		content := []rune{}
		read := 0

		for i := 0; i < count; i++ {
			r, num, ok := rf(test[read:])
			if !ok {
				return "", num, errors.New("Repeat is failed")
			}

			read += num
			content = append(content, r)
		}

		return string(content), read, nil
	}}
}

// String validates equal test stiring.
func String(needle string) *Parser {
	return &Parser{func(test []rune) (interface{}, int, error) {
		content := []rune{}
		read := 0
		for _, r := range needle {
			if read >= len(test) || r != test[read] {
				return string(content), read, errors.New("String is failed")
			}
			read++
			content = append(content, r)
		}
		return string(content), len(content), nil
	}}
}
