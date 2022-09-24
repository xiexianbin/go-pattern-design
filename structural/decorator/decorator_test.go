package decorator

import "testing"

func TestNewDecorator(t *testing.T) {
	d := NewDecorator()

	d.say()
}
