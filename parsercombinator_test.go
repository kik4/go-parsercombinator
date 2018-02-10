package parsercombinator

import "testing"

func TestAnyChar(t *testing.T) {
	t.Parallel()

	// success
	successCases := []struct {
		in, want string
	}{
		{"test", "t"},
	}
	for _, c := range successCases {
		got, err := AnyChar().Once().Parse(c.in)
		if err != nil || got != c.want {
			t.Error(err, c.in, got, c.want)
		}
	}

	// fail
	failCases := []string{
		"",
	}
	for _, c := range failCases {
		got, err := AnyChar().Once().Parse(c)
		if err == nil {
			t.Error(err, c, got)
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
		got, err := String(c.in1).Once().Parse(c.in2)
		if err != nil || got != c.want {
			t.Error(err, c.in1, c.in2, got, c.want)
		}
	}

	// fail
	failcases := []struct {
		in1, in2 string
	}{
		{"A", "test"},
		{"lo0000000000ng test-string", "short parse-string"},
	}
	for _, c := range failcases {
		got, err := String(c.in1).Once().Parse(c.in2)
		if err == nil {
			t.Error(err, c.in1, c.in2, got)
		}
	}
}
