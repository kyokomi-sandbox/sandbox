package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"io"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	goji.Get("/", func (c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello")
	})
	goji.Serve()
}
