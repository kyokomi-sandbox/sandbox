package main

import (
	"fmt"
	"os"
	"log"
	"errors"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_errbitclient"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

type errbitClient struct {
	n *_errbitclient.Notifier
}

var errbit errbitClient

func main() {

	type User struct {
		name string
		level int
		job string
	}

	errbit.n = _errbitclient.NewNotifier("c6b95be929ec97c2373d4398cda478c0")

	serve()
}


func serve() {
	fmt.Println("serve start")

	goji.Get("/", func (c web.C, w http.ResponseWriter, r *http.Request) {

		if err := errbit.n.Notify(errors.New("test"), r); err != nil {
			log.Println(err)
		}

		if f, err := os.Open("test.json"); err != nil {
			if err := errbit.n.Notify(err, r); err != nil {
				log.Println(err)
			}
		} else {
			f.Close()
		}
	})

	goji.Serve()
}
