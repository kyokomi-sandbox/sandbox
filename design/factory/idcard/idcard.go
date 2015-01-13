package idcard

import (
	"fmt"
	"github.com/kyokomi-sandbox/go-sandbox/design/factory/base"
)

type IDCard struct {
	owner string
}

func NewIDCard(owner string) base.Product {
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
var _ base.Product = (*IDCard)(nil)

type IDCardFactory struct {
	owners []string
}

func newIDCardFactory() base.Factory {
	i := IDCardFactory{}
	i.owners = make([]string, 0)
	return i
}

func NewIDCardFactory() base.AbstractFactory {
	i := newIDCardFactory()
	f := base.NewAbstractFactory(i)
	return f
}

func (i IDCardFactory) CreateProduct(owner string) base.Product {
	return NewIDCard(owner)
}

func (i IDCardFactory) RegisterProduct(p base.Product) {
	i.owners = append(i.owners, (p.(IDCard)).Owner())
}

func (i IDCardFactory) Owners() []string {
	return i.owners
}
var _ base.Factory = (*IDCardFactory)(nil)
