package mediator

import "testing"

func TestMediator(t *testing.T) {
	m := &Mediator{}

	t2 := Test{
		Mediator: m,
	}
	d := Dev{
		Mediator: m,
	}

	m.Dev = d
	m.Test = t2

	d.SendMessage("app x dev is done.")
	t2.SendMessage("begion to test app x...")
}
