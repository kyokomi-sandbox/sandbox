package main

import (
	termbox "github.com/nsf/termbox-go"
	"fmt"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	x := 0
	y := 0
	termbox.SetCursor(x, y)
	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			if ev.Key == termbox.KeyCtrlS {
				// Syncでコンソールが消えるっぽい
				x = 0
				y = 0
				termbox.Sync()
				termbox.SetCursor(x, y)
				termbox.Flush()
				break
			}

			// Ctrl+Xモード時のCtrl+Qでループを抜ける制御
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}

			if ev.Key == termbox.KeyEnter {
				x = 0
				y++
				termbox.SetCursor(x, y)
				termbox.Flush()
				break
			}

			// タイプした文字を出力
			msg := string(ev.Ch)
			for _, c := range msg {
				termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
				x++
			}

			// 画面端まで行ったら行を更新
			xx, _ := termbox.Size()
			if x == xx {
				x = 0
				y++
			}

			debugPrint(fmt.Sprintf("x = %d xx = %d", x, xx))
			termbox.SetCursor(x, y)
			termbox.Flush()
		case termbox.EventResize:
			// ターミナルのサイズ変更で呼ばれる
			termbox.Flush()
		case termbox.EventMouse:
			// ターミナルにタッチすると呼ばれる
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func debugPrint(a ...interface{}) {
	for argNum := 0; argNum < len(a); argNum++ {
		// always add spaces if we're doing Println
		arg := a[argNum]
		termboxLogPrint(fmt.Sprintf("%v", arg))
	}
}

func termboxLogPrint(msg string) {
	x := 0
	y := 20
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x++
	}
}
