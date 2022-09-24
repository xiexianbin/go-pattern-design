package singleton

import "testing"

func TestNew(t *testing.T) {
	s1 := New()
	s1["k1"] = "v1"
	s1["k2"] = "v2"

	s2 := New()
	if s2["k1"] != "v1" {
		t.Errorf("s2[k1] != v1")
	}
}
