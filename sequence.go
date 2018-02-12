package parsercombinator

import "errors"

// Once adapts parse rule once.
func (rf RuneFunc) Once() *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		str, num, succeeded := rf(test)
		if succeeded {
			return str, num, nil
		}
		return "", 0, errors.New("once is failed")
	}}
}

// AtLeastOnce adapts parse rule once or more.
func (rf RuneFunc) AtLeastOnce() *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		content := make([]byte, 0)
		read := 0
		str, num, succeed := rf(test[read:])
		for ; succeed; str, num, succeed = rf(test[read:]) {
			read += num
			content = append(content, str...)
		}

		if read == 0 {
			return "", 0, errors.New("AtLeastOnce is failed")
		}

		return string(content), read, nil
	}}
}

// Many adapts parse rule 0 or more times.
func (rf RuneFunc) Many() *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		content := make([]byte, 0)
		read := 0
		str, num, succeed := rf(test[read:])
		for ; succeed; str, num, succeed = rf(test[read:]) {
			read += num
			content = append(content, str...)
		}

		return string(content), read, nil
	}}
}

// Repeat adapts parse rule count times.
func (rf RuneFunc) Repeat(count int) *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		if count <= 0 {
			return "", 0, errors.New("Repeat needs 1 or more times")
		}

		content := make([]byte, 0)
		read := 0

		for i := 0; i < count; i++ {
			str, num, succeed := rf(test[read:])

			if !succeed {
				return "", 0, errors.New("Repeat is failed")
			}

			read += num
			content = append(content, str...)
		}

		return string(content), read, nil
	}}
}
