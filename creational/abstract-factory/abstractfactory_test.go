package abstract_factory

import "testing"

func TestConcreteFactory_CreateProduct(t *testing.T) {
	cf := &ConcreteFactory{}

	product := cf.CreateProduct()

	conProduct := product.(*ConcreateProduct)

	if conProduct.GetName() != "p1" {
		t.Errorf("abstractfactory test failed")
	}
}
