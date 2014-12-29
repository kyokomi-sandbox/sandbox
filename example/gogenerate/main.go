package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/kyokomi-sandbox/go-sandbox/example/gogenerate/fruit"
	"github.com/kyokomi-sandbox/go-sandbox/example/gogenerate/nepu"
)

//go:generate foobar
type Hoge struct {
	Name    string
	Num     int
	Message string
}

var memo = `
# gen手順

$ cd ./nepu
$ go get github.com/clipperhouse/gen
$ gen add github.com/clipperhouse/set
$ gen add github.com/clipperhouse/linkedlist
$ gen add github.com/clipperhouse/ring
$ gen

# stringer手順

$ go get golang.org/x/tools/cmd/stringer
$ cd fruit
$ go generate
`
func main() {
	fmt.Println("---------------------------------------")
	fmt.Println(memo)
	fmt.Println("---------------------------------------")

	var g nepu.NepuSlice
	g = append(g, nepu.Nepu{ID: 1000, Name: "Foo", GroupName: "Gopher's"})
	g = append(g, nepu.Nepu{ID: 1001, Name: "Bar", GroupName: "Gopher's"})
	g = append(g, nepu.Nepu{ID: 1002, Name: "Buzz", GroupName: "Gopher's"})
	g = append(g, nepu.Nepu{ID: 2000, Name: "Test1", GroupName: "Test's"})
	g = append(g, nepu.Nepu{ID: 2001, Name: "Test2", GroupName: "Test's"})
	g = append(g, nepu.Nepu{ID: 2002, Name: "Test3", GroupName: "Test's"})

	pp.Println("Member count      = ", g.Count(func(_ nepu.Nepu) bool { return true }))
	pp.Println("Member Bar?       = ", g.Where(func(m nepu.Nepu) bool { return m.Name == "Bar" }))
	pp.Println("Member Name       = ", g.GroupByString(func(m nepu.Nepu) string { return m.Name }))
	pp.Println("Member GroupName  = ", g.GroupByString(func(m nepu.Nepu) string { return m.GroupName }))

	fmt.Println("---------------------------------------")

	var fruit fruit.Fruit = fruit.Apple
	fmt.Println(fruit)
}
