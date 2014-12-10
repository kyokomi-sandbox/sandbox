package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"fmt"
)

func proxyExample() {
	srcAddr := ":3000"
	dstURL := "http://127.0.0.1:8000"
	dst, err := url.Parse(dstURL)
	if err != nil {
		log.Fatalln(err)
	}
	proxyHandler := httputil.NewSingleHostReverseProxy(dst)
	server := http.Server{
		Addr: srcAddr,
		Handler: proxyHandler,
	}
	log.Fatalln(server.ListenAndServe())
}

func proxyExample2() {
	goji.Handle("/*", hello)
	goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
