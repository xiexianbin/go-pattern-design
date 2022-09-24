package chain_of_responsibility

import "testing"

func TestFilterChain(t *testing.T) {
	fc := FilterChain{}
	fc.add(&AuthFilter{}).add(&LogFilter{})
	fc.Handle(Event{Name: "test Event"})
}
