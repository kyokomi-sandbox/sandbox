package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type PostMessage struct {
	Token string
}

const messageTemplate = `
HelloWorld!
> test1
>  aaaaaaaaaa
>
> test2
>  aaaaaaaaaaaa
`

func postSlack() {
	postUrl := "https://slack.com/api/chat.postMessage"
	client := &http.Client{}
	res, err := client.PostForm(postUrl, url.Values{
		"token":      {"aaaa-aaaaaaaaaaaa-aaaaaaaaaaaaa-aaaaaaaa-aaaaaaaaaaaa"},
		"channel":    {"aaaaaaaaaaaaa"},
		"text":       {messageTemplate},
		"username":   {"bot"},
		"icon_emoji": {":beers:"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
