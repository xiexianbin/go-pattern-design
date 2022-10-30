package iterator

type Visitor interface {
	GetName() string
}

type Teacher struct {
	Name string
}

func (t *Teacher) GetName() string {
	return t.Name
}

type IContainer interface {
	Add(v Visitor)
	Remove(index int)
}

type Container struct {
	list []Visitor
}

func (c *Container) Add(v Visitor) {
	c.list = append(c.list, v)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

type IIterator interface {
	Next() Visitor
	HasNext() bool
}

type Iterator struct {
	Container
	index int
}

func (i *Iterator) Next() Visitor {
	v := i.list[i.index]
	i.index += 1
	return v
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}

func NewIterator() *Iterator {
	return &Iterator{
		index:     0,
		Container: Container{},
	}
}
