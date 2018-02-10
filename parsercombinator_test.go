package parsercombinator

import "testing"

func TestAnyChar(t *testing.T) {
	s, error := AnyChar().Once().Parse("test")
	if error != nil || s != "t" {
		t.Error(error, s)
	}
}

func TestString(t *testing.T) {
	s, error := String("test").Once().Parse("test")
	if error != nil || s != "test" {
		t.Error(error, s)
	}
}
