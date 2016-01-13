package main

import (
	"github.com/zenazn/goji"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goji_/context"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goji_/foo"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {

	c := context.NewContext("goji sample")

	f := foo.NewService(c)

	goji.Use(Check1)
	goji.Use(Check2)
	goji.Get("/hello/:name", hello)
	goji.Get("/v2/hello/:name", f.HelloHandler)
	goji.Serve()
}

func Check1(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Check1")

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Check2(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Check2")

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
