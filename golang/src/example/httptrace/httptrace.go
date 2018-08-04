package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	"github.com/bitly/go-simplejson"
	"gopkg.in/go-pp/pp.v2"
)

func main() {
	getTrace()
}

func getTrace() {
	req, err := http.NewRequest("GET", "https://aphro-api.appspot.com/v1/articles", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set hooks
	trace := httptrace.ClientTrace{
		GetConn: func(h string) {
			log.Println("GetConn:", h)
		},
	}

	ctx := httptrace.WithClientTrace(context.Background(), &trace)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	j, err := simplejson.NewFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(j)
}
