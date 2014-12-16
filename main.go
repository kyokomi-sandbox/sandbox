package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"io"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/KyokomiSandbox/GoSandbox/stringer"
)

func main() {
	textExample()
	stringer.StringerExample()
}

func server() {
	goji.Get("/", func (c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello")
	})

	goji.Serve()
}
