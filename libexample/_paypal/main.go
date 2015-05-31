package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	ctx := context.Background()
	ctx = NewContext(ctx)

	ctx = WithAuthCallbackFunc(ctx, func(ctx context.Context, w http.ResponseWriter, r *http.Request, token *oauth2.Token) {
		fmt.Println(token)
	})
	ctx = WithAuthErrorFunc(ctx, func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	})

	kami.Context = ctx

	kami.Get("/login/paypal", LoginPayPal)
	kami.Get("/auth/paypal/callback", AuthPayPalCallback)

	log.Println("Starting server...")
	log.Println("GOMAXPROCS: ", cpus)
	kami.Serve()
}
