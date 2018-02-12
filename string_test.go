package parsercombinator

import "testing"

func TestString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, test, want1 string
		want2           int
		succeed         bool
	}{
		{"abc", "abc", "abc", 3, true},
		{"߷ÁÁ", "߷Á", "߷Á", 4, true},
		{"あいうえおabc", "あいうえおa", "あいうえおa", 16, true},
		{"🍺🍣🍺", "🍺🍣🍺", "🍺🍣🍺", 12, true},
		{"long input", "abc", "", 0, false},
		{"abc", "long test", "", 0, false},
		{"", "", "", 0, true},
	}
	for i, c := range cases {
		got, num, err := String(c.test).Text().Once().Parse(c.in)
		if !(got == c.want1 && num == c.want2 || !(c.succeed == (err == nil))) {
			t.Error(i, got, num, err, c)
		}
	}
}
