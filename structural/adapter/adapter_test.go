package adapter

import "testing"

func TestAdapter_Request(t *testing.T) {
	adaptee := Adaptee{}
	adaptee.doSth()

	adapter := &Adapter{
		Adaptee: adaptee,
	}
	adapter.Request()
}