package mediator

import "fmt"

type Department interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}

type Dev struct {
	Mediator *Mediator
}

func (d *Dev) SendMessage(message string) {
	d.Mediator.ForwardMessage(d, message)
}

func (d *Dev) ReceiveMessage(message string) {
	fmt.Printf("Dev Department receive message: %s\n", message)
}

type Test struct {
	Mediator *Mediator
}

func (t *Test) SendMessage(message string) {
	t.Mediator.ForwardMessage(t, message)
}

func (t *Test) ReceiveMessage(message string) {
	fmt.Printf("Test Department receive message: %s\n", message)
}

type Mediator struct {
	Dev
	Test
}

func (m *Mediator) ForwardMessage(d Department, message string) {
	switch d.(type) {
	case *Dev:
		m.Test.ReceiveMessage(message)
	case *Test:
		m.Dev.ReceiveMessage(message)
	default:
		fmt.Println("no impl Mediator...")
	}
}
