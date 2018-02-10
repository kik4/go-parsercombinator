package parsercombinator

import (
	"testing"
)

func TestSequence(t *testing.T) {
	t.Parallel()

	p := Sequence(
		func(args []interface{}) interface{} {
			return args[1].(string)
		},
		String("abc").Once(),
		AnyChar().Once(),
		String("def").Once(),
	)

	cases := []struct {
		in      string
		want1   interface{}
		want2   int
		succeed bool
	}{
		{"abc-def", "-", 7, true},
		{"abc---def", nil, 0, false},
	}
	for i, c := range cases {
		got, num, err := p.Parse(c.in)
		if !(got == c.want1 && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}

	// test panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Sequence(nil, String("panic").Once()).Parse("panic")
}

func TestOr(t *testing.T) {
	t.Parallel()

	p := Or(
		Digit().Once(),
		AnyChar().Once(),
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
