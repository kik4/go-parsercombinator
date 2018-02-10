package parsercombinator

import "testing"

func TestAnyChar(t *testing.T) {
	if AnyChar("test") != 't' {
		t.Error()
	}
}
