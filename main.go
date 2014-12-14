package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"io"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/unrolled/render"
	"github.com/kyokomi/GoSandbox/docomo"
)

func main() {
	textExample()
}

func server() {
	rd := render.New()
	d := docomo.NewDialogue()

	goji.Get("/", func (c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello")
	})
	goji.Get("/zatsudan/:message", func (c web.C, w http.ResponseWriter, r *http.Request) {
		res := d.Send(c.URLParams["message"])
		rd.JSON(w, http.StatusOK, string(res))
	})

	goji.Serve()
}
