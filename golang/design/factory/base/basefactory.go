package base

type Product interface {
	Use()
}

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(p Product)
}

type AbstractFactory struct {
	factory Factory
}

func NewAbstractFactory(f Factory) AbstractFactory {
	return AbstractFactory{factory: f}
}

func (f AbstractFactory) Create(owner string) Product {
	p := f.factory.CreateProduct(owner)
	f.factory.RegisterProduct(p)
	return p
}
