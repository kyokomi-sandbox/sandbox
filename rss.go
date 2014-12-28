package main

import (
	rss "github.com/jteeuwen/go-pkg-rss"
	"log"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"github.com/k0kubun/pp"
)

func keyInputRssExample() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			if ev.Key == termbox.KeyCtrlS {
				// Syncでコンソールが消えるっぽい
				termbox.Sync()
				break
			}

			// Ctrl+Xモード時のCtrl+Qでループを抜ける制御
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}

			if ev.Key == termbox.KeyEnter {
				fmt.Println()
				break
			}

			rssExample()

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

var feed = rss.New(5, true, chanHandler, itemHandler)

func rssExample() {
	fmt.Print("rssExample")
	if err := feed.Fetch("http://www.onsen.ag/blog/?feed=rss2&cat=23", nil); err != nil {
		log.Fatalln(err)
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Println("chanHandler")
	fmt.Println(len(newchannels), "new channel(s) in", feed.Url)
	for _, c := range newchannels {
		pp.Println(c)
	}
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Println("itemHandler")
	fmt.Println(len(newitems), "new item(s) in", ch.Title, "of", feed.Url)
	for _, item := range newitems {
		pp.Println(item)
	}
}
