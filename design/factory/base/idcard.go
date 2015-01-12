package base

import (
	"fmt"
)

type IDCard struct {
	owner string
}

func NewIDCard(owner string) Product {
	fmt.Println(owner + "のカードを作ります。")
	i := IDCard{}
	i.owner = owner
	return i
}

func (i IDCard) Use() {
	fmt.Println(i.owner + "のカードを使います。")
}

func (i IDCard) Owner() string {
	return i.owner
}
var _ Product = (*IDCard)(nil)

type IDCardFactory struct {
	owners []string
}

func newIDCardFactory() Factory {
	i := IDCardFactory{}
	i.owners = make([]string, 0)
	return i
}

func NewIDCardFactory() AbstractFactory {
	i := newIDCardFactory()
	f := NewAbstractFactory(i)
	return f
}

func (i IDCardFactory) createProduct(owner string) Product {
	return NewIDCard(owner)
}

func (i IDCardFactory) registerProduct(p Product) {
	i.owners = append(i.owners, (p.(IDCard)).Owner())
}

func (i IDCardFactory) Owners() []string {
	return i.owners
}
var _ Factory = (*IDCardFactory)(nil)
