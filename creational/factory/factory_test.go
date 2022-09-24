package factory

import (
	"reflect"
	"testing"
)

func TestShapeFactory_GetShape(t *testing.T) {
	var s Shape
	var err error
	f := ShapeFactory{}
	if s, err = f.GetShape(RectangleShape); err == nil {
		r := s.(*Rectangle)
		r.draw()
	} else {
		t.Fatal(err)
	}

	if s, err = f.GetShape(CircleShape); err == nil {
		c := reflect.ValueOf(s).Interface().(*Circle)
		c.draw()
	} else {
		t.Fatal(err)
	}
}
