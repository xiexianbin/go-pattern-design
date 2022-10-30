package state

import "fmt"

type State interface {
	Handle()
}

type Context struct {
	State State
	Value int
}

func NewContext(v int) *Context {
	c := &Context{
		Value: v,
	}
	c.Change(v)
	return c
}

func (c *Context) Handle() {
	c.State.Handle()
}

func (c *Context) Change(v int) {
	if v >= 60 {
		c.State = &StateA{}
	} else {
		c.State = &StateB{}
	}
}

type StateA struct{}

func (h *StateA) Handle() {
	fmt.Println("state A")
}

type StateB struct{}

func (h *StateB) Handle() {
	fmt.Println("state B")
}
