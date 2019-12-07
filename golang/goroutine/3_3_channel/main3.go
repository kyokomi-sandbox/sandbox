package main

import (
	"bytes"
	"fmt"
	"os"
)

// バッファ付きチャネルを使う例
func main() {
	var stdoutBuff bytes.Buffer         // インメモリのバッファ、出力が非決定的になるのを軽減
	defer stdoutBuff.WriteTo(os.Stdout) // プロセスが終了する前に確実にバッファが書き込まれるように

	intStream := make(chan int, 4) // キャパシティ4のバッファ付きチャネル
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
