package prototype

type Message struct {
	Context string
}

func (m *Message) Clone() *Message {
	r := *m
	return &r
}

func New(content string) *Message {
	return &Message{
		Context: content,
	}
}
