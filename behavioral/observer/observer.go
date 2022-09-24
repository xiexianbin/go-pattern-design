package observer

import "fmt"

type Event struct {
	Name string
}

type Observer interface {
	Notify(e Event)
}

type Dad struct {
}

func (d *Dad) Notify(e Event) {
	fmt.Printf("Dad get %s event, do something\n", e.Name)
}

type Mum struct {
}

func (m *Mum) Notify(e Event) {
	fmt.Printf("Mum get %s event, do something\n", e.Name)
}

type Child struct {
	Name      string
	Observers []Observer
}

func (c *Child) AddObserver(o Observer) *Child {
	c.Observers = append(c.Observers, o)
	return c
}

func (c *Child) DeleteObserver(o Observer) *Child {
	// ...
	return c
}

func (c *Child) WakeUp() {
	fmt.Printf("%s wakeup\n", c.Name)
	event := Event{
		Name: "wakeup",
	}
	for _, o := range c.Observers {
		o.Notify(event)
	}
}
