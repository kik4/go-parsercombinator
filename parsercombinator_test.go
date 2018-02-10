package parsercombinator

import (
	"testing"
	"unicode/utf8"
)

func TestSequence(t *testing.T) {
	t.Parallel()

	p := Sequence(
		String("abc").Once(),
		AnyChar().Once(),
		String("def").Once(),
	)

	// success
	{
		got, num, err := p.Parse("abc-def")
		if err != nil || got != "abc-def" || num != utf8.RuneCountInString("abc-def") {
			t.Error(err, got, num)
		}
	}

	// fail
	{
		got, num, err := p.Parse("abc--def")
		if err == nil {
			t.Error(err, got, num)
		}
	}
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
