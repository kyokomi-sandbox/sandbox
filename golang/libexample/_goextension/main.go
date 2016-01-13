package main

//go:generate go-extpoints

import (
	"reflect"

	"log"

	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goextension/extpoints"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goextension/plugins"
)

func main() {

	name := reflect.TypeOf(plugins.FooPlugin{}).Name()

	f, ok := extpoints.MyPlugins.Lookup(name)
	if !ok {
		log.Fatalf("not found %s\n", name)
	}

	f.Before("test")
	f.Run()
	f.After()
}
