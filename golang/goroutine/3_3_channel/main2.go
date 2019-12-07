package main

import (
	"fmt"
	"sync"
)

// 複数のゴルーチンを一度に開放する例
// sync.Condを使っても同じことができる
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // ここでチャネルから読み込めるようになるまでゴルーチンは待機
			fmt.Printf("%v has begun\n", 1)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin) // チャンネルを閉じる。これによってすべてのゴルーチンを同時に開放する
	wg.Wait()
}
