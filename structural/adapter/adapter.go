package adapter

import "fmt"

type Target interface {
	Request()
}

type ConcreteTarget struct{}

func (c *ConcreteTarget) Request() {
	fmt.Println("do something ...")
}

type Adaptee struct{}

func (a *Adaptee) doSth() {
	fmt.Println("Adaptee do something...")
}

type Adapter struct {
	Adaptee Adaptee
}

func (a *Adapter) Request() {
	a.Adaptee.doSth()
}
