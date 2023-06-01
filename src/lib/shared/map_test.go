package shared

import "testing"

// Test MapContainsKey with a valid key
func TestMapContainsKeyWithAValidKey(t *testing.T) {
	m := map[string]string{"a": "b", "c": "d"}
	key := "a"
	if !MapContainsKey(m, key) {
		t.Errorf("MapContainsKey was incorrect, got: %t, want: %t.", MapContainsKey(m, key), true)
	}
}
