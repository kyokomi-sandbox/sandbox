package main

import (
	"fmt"
)

// 明確にチャネルを所有するゴルーチンと、チャネルのブロックと閉じることを扱う消費者を作成する
func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() { // ゴルーチンより先にチャネルを生成したことに注意
			defer close(resultStream) // resultStreamを使ったあと確実に閉じられるように
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // 戻り値は読み込み専用チャネルとして返す
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
