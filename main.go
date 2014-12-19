package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KyokomiSandbox/GoSandbox/stringer"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	textExample()
	stringer.StringerExample()

	mqttExample()
}

func server() {
	goji.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello")
	})

	goji.Serve()
}
