package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/oauth2"
	"github.com/guregu/kami"
	"golang.org/x/net/context"
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

	kami.Context = ctx

	kami.Get("/login/paypal", LoginPayPal)
	kami.Get("/auth/paypal/callback", AuthPayPalCallback)

	log.Println("Starting server...")
	log.Println("GOMAXPROCS: ", cpus)
	kami.Serve()
}
