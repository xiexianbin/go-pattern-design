package command

import "fmt"

type Receiver interface {
	Execute()
}

type ReceiverA struct{}

func (a *ReceiverA) Execute() {
	fmt.Println("Receiver A Execute")
}

type ReceiverB struct{}

func (b *ReceiverB) Execute() {
	fmt.Println("Receiver B Execute")
}

type Command interface {
	Call()
}

type ConcreteCommandA struct {
	Receiver
}

func (ca *ConcreteCommandA) Call() {
	ca.Receiver.Execute()
}

type ConcreteCommandB struct {
	Receiver
}

func (cb *ConcreteCommandB) Call() {
	cb.Receiver.Execute()
}

type Invoker struct {
	cmds []Command
}

func (in *Invoker) AddCommand(c Command) {
	if in == nil {
		return
	}
	in.cmds = append(in.cmds, c)
}

func (in *Invoker) ExecuteCommand() {
	if in == nil || len(in.cmds) == 0 {
		return
	}
	for _, cmd := range in.cmds {
		cmd.Call()
	}
}
