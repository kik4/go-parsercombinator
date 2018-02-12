package parsercombinator

import "testing"

func TestOnce(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rule    RuneFunc
		in      string
		want1   string
		want2   int
		succeed bool
	}{
		{Letter(), "abc", "a", 1, true},
		{Letter(), "3ff", "", 0, false},
		{Letter(), "", "", 0, false},
	}
	for i, c := range cases {
		got, num, err := c.rule.Once().Parse(c.in)
		if !(got == c.want1 && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}
}

func TestAtLeastOnce(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rule    RuneFunc
		in      string
		want1   string
		want2   int
		succeed bool
	}{
		{Letter(), "abc", "abc", 3, true},
		{Letter(), "afa3ff", "afa", 3, true},
		{Letter(), "3ff", "", 0, false},
		{Letter(), "", "", 0, false},
		{String("abc").Text(), "abcabcabcd", "abcabcabc", 9, true},
	}
	for i, c := range cases {
		got, num, err := c.rule.AtLeastOnce().Parse(c.in)
		if !(got == c.want1 && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}
}

func TestMany(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rule    RuneFunc
		in      string
		want1   string
		want2   int
		succeed bool
	}{
		{Letter(), "abc", "abc", 3, true},
		{Letter(), "afa3ff", "afa", 3, true},
		{Letter(), "3ff", "", 0, true},
		{Letter(), "", "", 0, true},
		{String("abc").Text(), "abcabcabcd", "abcabcabc", 9, true},
	}
	for i, c := range cases {
		got, num, err := c.rule.Many().Parse(c.in)
		if !(got == c.want1 && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}
}

func TestRepeat(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rule    RuneFunc
		in      string
		count   int
		want1   string
		want2   int
		succeed bool
	}{
		{Letter(), "abc", 3, "abc", 3, true},
		{Letter(), "afa3ff", 3, "afa", 3, true},
		{Letter(), "3ff", 3, "", 0, false},
		{Letter(), "", 3, "", 0, false},
		{String("abc").Text(), "abcabcabcd", 2, "abcabc", 6, true},
		{Letter(), "abc", -3, "", 0, false},
	}
	for i, c := range cases {
		got, num, err := c.rule.Repeat(c.count).Parse(c.in)
		if !(got == c.want1 && num == c.want2) || !(c.succeed == (err == nil)) {
			t.Error(i, got, num, err, c)
		}
	}
}
