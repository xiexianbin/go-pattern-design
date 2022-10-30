package prototype

import (
	"fmt"
	"testing"
)

func TestMessage_Clone(t *testing.T) {
	m := New("test message")
	n := m.Clone()
	m.Context = "new message"
	fmt.Println(m.Context)
	fmt.Println(n.Context)
}
