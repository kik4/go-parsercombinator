package parsercombinator

import "errors"

// Once adapts parse rule once.
func (rf RuleFunc) Once() *Parser {
	return &Parser{func(test string) (string, int, error) {
		str, num, succeeded := rf(test)
		if succeeded {
			return str, num, nil
		}
		return "", 0, errors.New("once is failed")
	}}
}

// AtLeastOnce adapts parse rule once or more.
func (rf RuleFunc) AtLeastOnce() *Parser {
	return &Parser{func(test string) (string, int, error) {
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
