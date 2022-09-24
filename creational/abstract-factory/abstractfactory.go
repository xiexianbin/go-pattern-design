package abstract_factory

type AbstractFactory interface {
	CreateProduct() AbstractProduct
}

type AbstractProduct interface {
	GetName() string
}

type ConcreateProduct struct {
	Name string
}

func (c *ConcreateProduct) GetName() string {
	return c.Name
}

type ConcreteFactory struct{}

func (c *ConcreteFactory) CreateProduct() AbstractProduct {
	return &ConcreateProduct{Name: "p1"}
}
