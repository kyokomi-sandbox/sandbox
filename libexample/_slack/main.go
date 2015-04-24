package main

import (
	"flag"

	"github.com/kyokomi/slack"
)

func main() {
	token := flag.String("token", "", "")
	room := flag.String("room", "#random", "")
	message := flag.String("message", "hoge", "")
	flag.Parse()
	if *token == "" || *room == "" || *message == "" {
		panic("arg error")
	}
	slack.New(*token).PostMessage(*room, *message, slack.NewPostMessageParameters())
}
