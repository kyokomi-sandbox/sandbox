package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/k0kubun/pp"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {

	dstURL := "http://127.0.0.1:8080"
	dst, err := url.Parse(dstURL)
	if err != nil {
		log.Fatalln(err)
	}
	proxyHandler := httputil.NewSingleHostReverseProxy(dst)

	goji.Use(func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			pp.Println(r.Header)
			pp.Println(r.Cookies())
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	})

	goji.Handle("/*", proxyHandler)
	goji.Serve()
}

func defaultHttp() {
	srcAddr := ":3000"
	dstURL := "http://127.0.0.1:8080"
	dst, err := url.Parse(dstURL)
	if err != nil {
		log.Fatalln(err)
	}
	proxyHandler := httputil.NewSingleHostReverseProxy(dst)
	server := http.Server{
		Addr:    srcAddr,
		Handler: proxyHandler,
	}
	log.Fatalln(server.ListenAndServe())
}
