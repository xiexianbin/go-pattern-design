package state

import "testing"

func TestNewContext(t *testing.T) {
	c := NewContext(99)
	if _, ok := c.State.(*StateA); !ok {
		t.Error("state err")
	}

	c.Change(50)
	if _, ok := c.State.(*StateB); !ok {
		t.Error("state err")
	}
}
