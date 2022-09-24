package strategy

import "fmt"

type Strategy interface {
	Charge()
}

type St1 struct{}

func (s *St1) Charge() {
	fmt.Println("st1...")
}

type St2 struct{}

func (s *St2) Charge() {
	fmt.Println("st2...")
}

type Context struct {
	Strategy Strategy
}

func (c *Context) DoCharge() {
	c.Strategy.Charge()
}
