package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	// Getが呼び出されたときに、まずプール内に使用可能なインスタンスがあるか確認し、あるいは呼出もとにそれを返す
	// もしなければ、Newメンバー変数を呼び出し、新しいインスタンスを作成する
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}
