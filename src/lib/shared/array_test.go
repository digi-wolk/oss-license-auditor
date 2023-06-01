package shared

import "testing"

// Test ArrayContains with a valid string
func TestArrayContainsWithAValidString(t *testing.T) {
	s := []string{"a", "b", "c"}
	str := "b"
	if !ArrayContains(s, str) {
		t.Errorf("ArrayContains was incorrect, got: %t, want: %t.", ArrayContains(s, str), true)
	}
}
