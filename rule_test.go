package parsercombinator

import "testing"

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
	for _, c := range cases {
		got, num, succeed := Letter()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(got, num, succeed, c)
		}
	}
}
