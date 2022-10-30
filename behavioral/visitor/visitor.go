package visitor

import "fmt"

type Element interface {
	doSth()
	accept(visitor IVisitor)
}

type ConcreteElement1 struct{}

func (c1 *ConcreteElement1) doSth() {
	fmt.Println("c1 doSth ...")
}

func (c1 *ConcreteElement1) accept(visitor IVisitor) {
	visitor.visit1(c1)
}

type ConcreteElement2 struct{}

func (c2 *ConcreteElement2) doSth() {
	fmt.Println("c2 doSth ...")
}

func (c2 *ConcreteElement2) accept(visitor IVisitor) {
	visitor.visit2(c2)
}

type IVisitor interface {
	visit1(c1 *ConcreteElement1)
	visit2(c2 *ConcreteElement2)
}

type Visitor struct{}

func (v *Visitor) visit1(c1 *ConcreteElement1) {
	c1.doSth()
}

func (v *Visitor) visit2(c2 *ConcreteElement2) {
	c2.doSth()
}

type ObjectStruture struct{}

func (o *ObjectStruture) createElement() {
	v := new(Visitor)
	c1 := ConcreteElement1{}
	c1.accept(v)
	c2 := ConcreteElement2{}
	c2.accept(v)

	//v.visit1(&c1)
	//v.visit2(&c2)
}
