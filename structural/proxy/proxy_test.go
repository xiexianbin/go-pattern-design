package proxy

import "testing"

func TestProxy(t *testing.T) {
	c := &Company{
		action: "do something ...",
	}

	p := new(ProxyCompany)
	p.c = c

	p.do()
}
