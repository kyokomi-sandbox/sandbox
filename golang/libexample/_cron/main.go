package main

import (
	"fmt"

	"errors"
	"net/http"

	"github.com/guregu/kami"
	"github.com/robfig/cron"
	"golang.org/x/net/context"
)

func main() {
	c := cron.New()
	c.AddFunc("5 * * * * *", Hoge)   // 毎分5秒
	c.AddFunc("*/5 * * * * *", Fuga) // 5秒毎

	c.Start()
	isRunning := true

	kami.Get("/", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	kami.Get("/start", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		if !isRunning {
			c.Start()
			isRunning = true
		}
		w.Write([]byte("START OK"))
	})
	kami.Get("/stop", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		if isRunning {
			c.Stop()
			isRunning = false
		}
		w.Write([]byte("STOP OK"))
	})

	defer func() {
		if isRunning {
			c.Stop()
			isRunning = false
		}
	}()

	kami.Serve()
}

func Hoge() {
	fmt.Println("Every hour on the half hour")

	defer func() {
		fmt.Println(recover())
	}()

	panic(errors.New("Hoge errorだよ"))
}

func Fuga() {
	fmt.Println("Every hour on the half hour")

	defer func() {
		fmt.Println(recover())
	}()

	panic(errors.New("Fuga errorだよ"))
}
