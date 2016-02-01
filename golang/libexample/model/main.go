package main

import (
	"fmt"

	"gopkg.in/jeevatkm/go-model.v0"
)

func main() {
	type Hoge struct {
		ID       string
		Name     *string
		Number   *int
		Children []*Hoge
	}

	hoge := Hoge{}
	hoge.ID = "hoge"
	name := "hogehoge"
	hoge.Name = &name
	number := 110000
	hoge.Number = &number
	hoge.Children = []*Hoge{}

	fuga := model.Clone(&hoge)

	fmt.Printf("%s(%v) %d(%v)\n",
		*hoge.Name, hoge.Name, *hoge.Number, hoge.Number)
	fmt.Printf("%s(%v) %d(%v)\n",
		*fuga.(*Hoge).Name, fuga.(*Hoge).Name, *fuga.(*Hoge).Number, fuga.(*Hoge).Number)

	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))
	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))
	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))
	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))
	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))
	hoge.Children = append(hoge.Children, model.Clone(&hoge).(*Hoge))

	hoge2 := model.Clone(&hoge)

	fmt.Printf("%v\n",
		hoge.Children)
	fmt.Printf("%v\n",
		hoge2.(*Hoge).Children)
}
