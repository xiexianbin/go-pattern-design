package command

import "testing"

func TestInvoker_ExecuteCommand(t *testing.T) {
	ra := &ReceiverA{}
	rb := &ReceiverB{}

	ca := &ConcreteCommandA{
		Receiver: ra,
	}
	cb := &ConcreteCommandB{
		Receiver: rb,
	}

	i := Invoker{}
	i.AddCommand(ca)
	i.AddCommand(cb)
	i.ExecuteCommand()
}
