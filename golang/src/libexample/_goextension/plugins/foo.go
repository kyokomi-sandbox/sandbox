package plugins

import (
	"fmt"

	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goextension/extpoints"
)

func init() {
	// FooPlugin Pluginを登録
	extpoints.MyPlugins.Register(new(FooPlugin), "")
}

type FooPlugin struct {
	name string
}

func (f *FooPlugin) Before(name string) {
	f.name = name
}

func (f *FooPlugin) Run() {
	fmt.Printf("***** %s *****\n", f.name)
}

func (f *FooPlugin) After() {
	fmt.Printf("%s end\n", f.name)
}

func (f FooPlugin) String() string {
	return "fooPlugin"
}
