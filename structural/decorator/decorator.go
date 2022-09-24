package decorator

import "fmt"

type Component interface {
	say()
}

type ConcreateComponent struct{}

func (c *ConcreateComponent) say() {
	fmt.Println("say something ...")
}

type Decorator struct {
	c *ConcreateComponent
}

func (d *Decorator) say() {
	fmt.Println("decorator say things ...")
	d.c.say()
}

func NewDecorator() *Decorator {
	return &Decorator{
		c: &ConcreateComponent{},
	}
}
