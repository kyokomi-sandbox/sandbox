package main

import (
	"fmt"
	"github.com/kyokomi-sandbox/go-sandbox/design/factory/base"
)

/**
気合で実装したけど、実用的じゃない。。。
base.FactoryとFactoryが別物になるのでキャストできない
 */
func main() {
	fmt.Println("hello")

	f := base.NewIDCardFactory()
	card1 := f.Create("結城友奈")
	card2 := f.Create("暁美ほむら")
	card1.Use()
	card2.Use()
}
