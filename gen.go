package main

import (
	"github.com/KyokomiSandbox/GoSandbox/generate"
	"github.com/k0kubun/pp"
)

func genExample() {
	var g generate.MySliceTypeSlice
	g = append(g, generate.MySliceType{ID: 1000, Name: "Foo", GroupName: "Gopher's"})
	g = append(g, generate.MySliceType{ID: 1001, Name: "Bar", GroupName: "Gopher's"})
	g = append(g, generate.MySliceType{ID: 1002, Name: "Buzz", GroupName: "Gopher's"})
	g = append(g, generate.MySliceType{ID: 2000, Name: "Test1", GroupName: "Test's"})
	g = append(g, generate.MySliceType{ID: 2001, Name: "Test2", GroupName: "Test's"})
	g = append(g, generate.MySliceType{ID: 2002, Name: "Test3", GroupName: "Test's"})

	pp.Println("Member count      = ", g.Count(func(_ generate.MySliceType) bool { return true }))
	pp.Println("Member Bar?       = ", g.Where(func(m generate.MySliceType) bool { return m.Name == "Bar" }))
	pp.Println("Member Name       = ", g.GroupByString(func(m generate.MySliceType) string { return m.Name }))
	pp.Println("Member GroupName  = ", g.GroupByString(func(m generate.MySliceType) string { return m.GroupName }))
}
