package decorator

import "testing"

func Double(n int) int {
	return n * 2
}

func TestLogDecorate(t *testing.T) {
	f := LogDecorate(Double)
	if f(5) != 10 {
		t.Errorf("5*2 != 10")
	}
}
