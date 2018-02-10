package parsercombinator

import (
	"testing"
	"unicode/utf8"
)

func TestAnyChar(t *testing.T) {
	t.Parallel()

	// success
	successCases := []struct {
		in, want string
	}{
		{"test", "t"},
	}
	for _, c := range successCases {
		got, num, err := AnyChar().Once().Parse(c.in)
		if err != nil || got != c.want || num != 1 {
			t.Error(err, c.in, got, c.want)
		}
	}

	// fail
	failCases := []string{
		"",
	}
	for _, c := range failCases {
		got, num, err := AnyChar().Once().Parse(c)
		if err == nil {
			t.Error(err, c, got, num)
		}
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	// success
	successCases := []struct {
		in1, in2, want string
	}{
		{"", "", ""},
		{"test", "test", "test"},
	}
	for _, c := range successCases {
		got, num, err := String(c.in1).Once().Parse(c.in2)
		if err != nil || got != c.want || num != utf8.RuneCountInString(got) {
			t.Error(err, c.in1, c.in2, got, c.want)
		}
	}

	// fail
	failcases := []struct {
		in1, in2 string
	}{
		{"A", "test"},
		{"loooooooooooooong test-string", "short parse-string"},
	}
	for _, c := range failcases {
		got, _, err := String(c.in1).Once().Parse(c.in2)
		if err == nil {
			t.Error(err, c.in1, c.in2, got)
		}
	}
}

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
