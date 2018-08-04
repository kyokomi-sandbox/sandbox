package main

import (
	"flag"
	"time"

	"log"

	"fmt"

	"github.com/kyokomi/slack"
)

func main() {
	log.SetFlags(log.Llongfile)

	token := flag.String("token", "", "")
	room := flag.String("room", "#general", "")
	message := flag.String("message", "hoge", "")
	flag.Parse()
	if *token == "" || *room == "" || *message == "" {
		panic("arg error")
	}
	chSender := make(chan slack.OutgoingMessage)
	chReceiver := make(chan slack.SlackEvent)

	api := slack.New(*token)
	api.SetDebug(true)
	wsAPI, err := api.StartRTM("", "http://example.com")
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	go wsAPI.HandleIncomingEvents(chReceiver)
	go wsAPI.Keepalive(20 * time.Second)
	go func(wsAPI *slack.SlackWS, chSender chan slack.OutgoingMessage) {
		for {
			select {
			case msg := <-chSender:
				wsAPI.SendMessage(&msg)
			}
		}
	}(wsAPI, chSender)

	for {
		select {
		case msg := <-chReceiver:
			fmt.Print("Event Received: ")
			switch msg.Data.(type) {
			case slack.HelloEvent:
				// TODO: デフォルトChannelに何か投げたい
			case *slack.MessageEvent:
				a := msg.Data.(*slack.MessageEvent)
				fmt.Printf("Message: %v\n", a)
				wsAPI.SendMessage(wsAPI.NewOutgoingMessage(a.Text + "だってさ!", a.ChannelId))
			case *slack.PresenceChangeEvent:
				a := msg.Data.(*slack.PresenceChangeEvent)
				fmt.Printf("Presence Change: %v\n", a)
			case slack.LatencyReport:
				a := msg.Data.(slack.LatencyReport)
				fmt.Printf("Current latency: %v\n", a.Value)
			case *slack.SlackWSError:
				error := msg.Data.(*slack.SlackWSError)
				fmt.Printf("Error: %d - %s\n", error.Code, error.Msg)
			default:
				fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
