package main

import (
	"fmt"
)

// 拘束パターン
func main() {
	chanOwner := func() <-chan int {
		// チャネルをchanOwner関数のレキシカルスコープ内で初期化する
		// これによってresultsチャンネルへの書き込みができるスコープを制限
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// intのチャネルの読み込み専用のコピーを受け取る
	// 読み込み権限のみが必要であることを宣言することでこの関数内でのチャネルに対する操作を読み込み専用に拘束する
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	// チャネルへの読み込み権限を受け取って、消費者にわたす
	// 消費者は読み込み以外は何もしない
	results := chanOwner()
	consumer(results)
}
