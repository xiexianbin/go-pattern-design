package facade

import "fmt"

type A struct{}

func (a *A) Open() {
	fmt.Println("Open Door ...")
}

type B struct{}

func (b *B) Pack() {
	fmt.Println("Pack ...")
}

type C struct{}

func (c *C) Close() {
	fmt.Println("Close Door ...")
}

type Facade struct {
	a *A
	b *B
	c *C
}

func NewFacade() *Facade {
	return &Facade{
		a: &A{},
		b: &B{},
		c: &C{},
	}
}

func (f *Facade) Do() {
	f.a.Open()
	f.b.Pack()
	f.c.Close()
}
