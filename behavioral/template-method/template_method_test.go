package template_method

import "testing"

func TestTm(t *testing.T) {
	d := &Dog{Name: "d1"}
	d.activities()

	c := &Cat{Name: "tom"}
	c.activities()
}
