package parsercombinator

import (
	"testing"
)

/*
func TestSequence(t *testing.T) {
	t.Parallel()

	parsers := []*Parser{
		String("abc").Once(),
		AnyRune().Once(),
		String("def").Once(),
	}

	p1 := Sequence(
		parsers,
		func(args []interface{}) interface{} {
			return args[1].(string)
		},
	)
	p2 := Sequence(
		parsers,
		nil,
	)
	p3 := Sequence(
		[]*Parser{
			String("テスト").Once(),
			AnyRune().Once(),
			String("🍣").Once(),
			p1,
		},
		func(args []interface{}) interface{} {
			return args[1].(string) + args[3].(string)
		},
	)

	cases := []struct {
		parser  *Parser
		in      string
		want1   interface{}
		want2   int
		succeed bool
	}{
		{p1, "abc-def", "-", 7, true},
		{p1, "abc---def", nil, 0, false},
		{p2, "abcAdef", []interface{}{"abc", "A", "def"}, 7, true},
		{p3, "テスト🍺🍣abc❤def", "🍺❤", 26, true},
	}
	for i, c := range cases {
		got, num, err := c.parser.Parse(c.in)
		if !(reflect.DeepEqual(got, c.want1) && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}
}
*/
func TestOr(t *testing.T) {
	t.Parallel()

	p := Or(
		Digit().Once(),
		AnyRune().Once(),
	)

	// success
	successCases := []struct {
		in, want string
	}{
		{"0", "0"},
		{"b", "b"},
	}
	for _, c := range successCases {
		got, num, err := p.Parse(c.in)
		if err != nil || got != c.want || num != 1 {
			t.Error(err, got, num)
		}
	}

	// fail
	failcases := []struct {
		in string
	}{
		{""},
	}
	for _, c := range failcases {
		got, _, err := p.Parse(c.in)
		if err == nil {
			t.Error(err, c.in, got)
		}
	}
}
