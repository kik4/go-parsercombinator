package parsercombinator

import "errors"

// String validates equal test stiring.
func String(needle string) *Parser {
	return &Parser{func(test string) (interface{}, int, error) {
		length := len(needle)
		if len(test) < length {
			return "", 0, errors.New("String parse filed")
		}
		if test[:length] != needle {
			return "", 0, errors.New("String parse filed")
		}
		return test[:length], length, nil
	}}
}

// Text returns string value RuneFunc
func (p *Parser) Text() RuneFunc {
	return func(test string) (string, int, bool) {
		val, num, err := p.f(test)

		if err != nil {
			return "", 0, false
		}

		return val.(string), num, true
	}
}
