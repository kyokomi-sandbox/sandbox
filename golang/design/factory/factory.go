package main

import (
	"fmt"
	"github.com/kyokomi-sandbox/go-sandbox/design/factory/idcard"
	"github.com/kyokomi-sandbox/go-sandbox/design/factory/base"
)

func main() {
	fmt.Println("hello")

	var f base.AbstractFactory
	f = idcard.NewIDCardFactory()

	card1 := f.Create("結城友奈")
	card2 := f.Create("暁美ほむら")
	card1.Use()
	card2.Use()

	var s base.AbstractFactory
	s = NewSampleFactory()
	s1 := s.Create("aa")
	s2 := s.Create("bb")
	s1.Use()
	s2.Use()
}

type SampleProduct struct {
	name string
}

func (s SampleProduct) Use() {
	fmt.Println("sample use ", s.name)
}

// SampleFactory is Factory interface
type SampleFactory struct {

}

func NewSampleFactory() base.AbstractFactory {
	s := SampleFactory{}
	f := base.NewAbstractFactory(s)
	return f
}

func (s SampleFactory) CreateProduct(owner string) base.Product {
	return SampleProduct{name: owner}
}

func (s SampleFactory) RegisterProduct(p base.Product) {
	fmt.Println("sample register ", p.(SampleProduct).name)
}
