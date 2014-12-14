package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"io"
	"github.com/zenazn/goji/web"
	"net/http"
	docomo "github.com/kyokomi/go-docomo"
	"log"
)

func main() {
	textExample()

	d := docomo.New()
	res, err := d.SendImage("/Users/kyokomi/Downloads/menu_shop.png")
//	res, err := d.SendImage("/Users/kyokomi/src/github.com/kyokomi/GoSandbox/image2.png")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(res))
}

func server() {
	goji.Get("/", func (c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello")
	})

	goji.Serve()
}
