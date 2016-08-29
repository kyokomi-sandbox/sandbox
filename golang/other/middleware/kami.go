// Our webserver
package main

import (
	"fmt"
	"net/http"

	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

func greetMiddleware(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	fmt.Println("hoge")
	return ctx
}

func greet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := kami.Param(ctx, "name")
	fmt.Fprintf(w, "%s, %s!", "hello", name)
}

func main() {
	ctx := context.Background()
	kami.Context = ctx // set our "god context", the base context for all requests

	kami.Use("/", greetMiddleware)  // use this middleware for paths under /hello/
	kami.Get("/hello/:name", greet) // add a GET handler with a parameter in the URL
	kami.Serve()                    // gracefully serve with support for einhorn and systemd
}
