package chain_of_responsibility

import "fmt"

type Filter interface {
	Handle(event Event)
	//Handle(event Event, fc FilterChain)
}

type Event struct {
	Name string
}

type AuthFilter struct{}

func (f *AuthFilter) Handle(event Event) {
	fmt.Printf("do AuthFilter for %s\n", event.Name)
}

type LogFilter struct{}

func (f *LogFilter) Handle(event Event) {
	fmt.Printf("do LogFilter for %s\n", event.Name)
}

type FilterChain struct {
	Filters []*Filter
	// index int // which filter has run handle
}

func (fc *FilterChain) add(f Filter) *FilterChain {
	fc.Filters = append(fc.Filters, &f)
	return fc
}

func (fc *FilterChain) Handle(event Event) {
	for _, f := range fc.Filters {
		(*f).Handle(event)
	}
}
