package parsercombinator

import "testing"

func TestAnyChar(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"abc", "a", 1, true},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := AnyChar()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, test, want1 string
		want2           int
		want3           bool
	}{
		{"abc", "abc", "abc", 3, true},
		{"long input", "abc", "", 0, false},
		{"abc", "long test", "", 0, false},
		{"", "", "", 0, true},
	}
	for i, c := range cases {
		got, num, succeed := String(c.test)(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestDigit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"000", "0", 1, true},
		{"abc", "", 0, false},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := Digit()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestLetter(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"abc", "a", 1, true},
		{"000", "", 0, false},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := Letter()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}
