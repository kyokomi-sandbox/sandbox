package main

import (
	"fmt"
	"sync"
)

// はじめに確保した4LBとあろけーしょｎ4KBの合計8KBですむ
// Poolが便利なその他の場面としては、
// 可能な限り素早く実行しなければならない操作のためにアロケート済みのオブジェクトを暖気する状況
func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// プールに4KB確保する
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// 何かする
			// メモリに対して素早い処理が行われること
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
