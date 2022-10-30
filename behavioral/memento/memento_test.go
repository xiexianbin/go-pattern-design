package memento

import "testing"

func TestCaretaker(t *testing.T) {
	o := new(Originator)
	o.SetState("init state ...")

	// manage Caretaker
	c := &Caretaker{}
	m := c.CreateMemento(*o)
	if m.GetState() != o.state {
		t.Errorf("memento state not match err")
	}

	// update state
	o.SetState("new state ...")
	if m.GetState() == o.state {
		t.Errorf("change state err")
	}

	// recover state
	*o = c.RecoverOriginator(m)
	t.Log(o.GetState())
}