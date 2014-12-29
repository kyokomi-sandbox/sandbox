package main

import (
	rss "github.com/jteeuwen/go-pkg-rss"
	"log"
	"fmt"
	"github.com/k0kubun/pp"
)

func main () {
	rssExample()
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
