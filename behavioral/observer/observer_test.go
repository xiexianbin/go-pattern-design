package observer

import "testing"

func TestChild_WakeUp(t *testing.T) {
	c := Child{Name: "abc"}
	c.AddObserver(new(Dad)).AddObserver(new(Mum))
	c.WakeUp()
}
