package base

type Product interface {
	Use()
}

type Factory interface {
	createProduct(owner string) Product
	registerProduct(p Product)
}

type AbstractFactory struct {
	factory Factory
}

func NewAbstractFactory(f Factory) AbstractFactory {
	return AbstractFactory{factory: f}
}

func (f AbstractFactory) Create(owner string) Product {
	p := f.factory.createProduct(owner)
	f.factory.registerProduct(p)
	return p
}
