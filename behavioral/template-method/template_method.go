package template_method

import "fmt"

type Animal interface {
	eat()
	sleep()
	otherThing()
	activities()
}

type Dog struct {
	Name string
}

func (d *Dog) eat() {
	fmt.Printf("%s eat\n", d.Name)
}

func (d *Dog) sleep() {
	fmt.Printf("%s sleep\n", d.Name)
}

func (d *Dog) otherThing() {
	fmt.Printf("%s wangwang\n", d.Name)
}

func (d *Dog) activities() {
	d.eat()
	d.sleep()
	d.otherThing()
}

type Cat struct {
	Name string
}

func (c *Cat) eat() {
	fmt.Printf("%s eat\n", c.Name)
}

func (c *Cat) sleep() {
	fmt.Printf("%s sleep\n", c.Name)
}

func (c *Cat) otherThing() {
	fmt.Printf("%s miaomiao\n", c.Name)
}

func (c *Cat) activities() {
	c.eat()
	c.sleep()
	c.otherThing()
}
