package factory

import (
	"errors"
	"fmt"
)

type Type int

const (
	RectangleShape Type = 1 << iota
	CircleShape
)

type Shape interface {
	draw()
}

type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("run Rectangle:draw() ...")
}

type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("run Circle:draw() ...")
}

type ShapeFactory struct{}

func (s *ShapeFactory) GetShape(t Type) (Shape, error) {
	switch t {
	case RectangleShape:
		r := new(Rectangle)
		return r, nil
	case CircleShape:
		return &Circle{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("un-support type %#v", t))
	}
}
