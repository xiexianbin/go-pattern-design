package proxy

import "fmt"

type ICompany interface {
	do()
}

type Company struct {
	action string
}

func (c *Company) do() {
	fmt.Println(c.action)
}

type ProxyCompany struct {
	c *Company
}

func (p *ProxyCompany) do() {
	p.c.do()
}
