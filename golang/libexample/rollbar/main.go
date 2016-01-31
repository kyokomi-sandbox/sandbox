package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/guregu/kami"
	"github.com/stvp/rollbar"
	"github.com/unrolled/render"
	"golang.org/x/net/context"
)

var renderer = render.New(render.Options{})

func init() {
	log.SetFlags(log.Llongfile)
	var projectID int64
	var apiKey string
	flag.Int64Var(&projectID, "p", 0, "projectID")
	flag.StringVar(&apiKey, "a", "", "apiKey")
	flag.Parse()

	rollbar.Token = ""
}

func main() {
	ctx := context.Background()
	kami.Context = ctx

	kami.PanicHandler = PanicHandler
	kami.Get("/", index)

	rollbar.Message(rollbar.INFO, "test")

	kami.Serve()
}

// PanicHandler send errbit 500 errors
func PanicHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	panicSendAirbrake(ctx, req)
	renderer.JSON(w, 500, "TODO: 500")
}

func panicSendAirbrake(ctx context.Context, req *http.Request) {
	exception := kami.Exception(ctx)

	log.Println("ERROR:", exception)
	log.Println("ERROR:", string(debug.Stack()))

	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR:", err)
		}
	}()

	rollbar.RequestError(rollbar.ERR, req, fmt.Errorf("%s", exception))
}

/*
1. goroutineで投げるとbacktraceがおえなくなる
2. 同じbacktraceだと一つのエラーに集約される
*/

func index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("Error: %s", "hogehoge")

	sendError1Func := func(err error) {
		rollbar.RequestError(rollbar.ERR, r, fmt.Errorf("sendError1 : %s", err))
	}
	sendError2Func := func(err error) {
		rollbar.RequestError(rollbar.ERR, r, err)
	}

	// そのままerr
	rollbar.RequestError(rollbar.ERR, r, err)

	// fmt.Errorfでラップ
	rollbar.RequestError(rollbar.ERR, r, fmt.Errorf("Error: %s", err.Error()))

	sendError1Func(err)

	sendError2Func(fmt.Errorf("sendError2 : %s", err))

	panic("hogehoge")
}
