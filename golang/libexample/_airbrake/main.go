package main

import (
	"log"
	"net/http"
	"runtime/debug"

	"fmt"

	"flag"

	"github.com/guregu/kami"
	"github.com/unrolled/render"
	"golang.org/x/net/context"
	"gopkg.in/airbrake/gobrake.v1"
)

var renderer = render.New(render.Options{})
var airbrake *gobrake.Notifier

func init() {
	var projectID int64
	var apiKey string
	flag.Int64Var(&projectID, "p", 0, "projectID")
	flag.StringVar(&apiKey, "a", "", "apiKey")
	flag.Parse()
	airbrake = gobrake.NewNotifier(projectID, apiKey)
}

func main() {
	ctx := context.Background()
	kami.Context = ctx

	kami.PanicHandler = PanicHandler
	kami.Get("/", index)

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

	err := airbrake.Notify(exception, req)
	if err != nil {
		log.Println("ERROR:", "airbrake send error: %s (request error: %s)", err, exception)
	}
}

/*
1. goroutineで投げるとbacktraceがおえなくなる
2. 同じbacktraceだと一つのエラーに集約される
*/

func index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("Error: %s", "hogehoge")

	sendError1Func := func(err error) {
		airbrake.Notify(fmt.Errorf("sendError1 : %s", err), r)
	}
	sendError2Func := func(err error) {
		airbrake.Notify(err, r)
	}

	// そのままerr
	airbrake.Notify(err, r)

	// fmt.Errorfでラップ
	airbrake.Notify(fmt.Errorf("Error: %s", err.Error()), r)

	sendError1Func(err)

	sendError2Func(fmt.Errorf("sendError2 : %s", err))

	panic("hogehoge")
}
