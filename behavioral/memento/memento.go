package memento

type IState interface {
	GetState() string
	SetState(state string)
}

type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

type Memento struct {
	Originator
}

func (m *Memento) GetState() string {
	return m.Originator.state
}

func (m *Memento) SetState(originator Originator) {
	m.Originator = originator
}

type Caretaker struct{}

func (c *Caretaker) CreateMemento(originator Originator) Memento {
	return Memento{
		Originator: originator,
	}
}

func (c *Caretaker) RecoverOriginator(memento Memento) Originator {
	return memento.Originator
}
