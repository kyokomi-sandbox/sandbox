package main

import (
	"fmt"

	"github.com/robfig/cron"
	"time"
	"github.com/juju/errors"
)

func main() {
	c := cron.New()
	c.AddFunc("5 * * * * *", Hoge) // 毎分5秒
	c.AddFunc("*/5 * * * * *", Fuga) // 5秒毎
	c.Start()
	defer c.Stop()

	time.Sleep(1 * time.Minute)
}

func Hoge() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	fmt.Println("Every hour on the half hour")

	panic(errors.New("Hoge errorだよ"))
}

func Fuga() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	fmt.Println("Every hour on the half hour")

	panic(errors.New("Fuga errorだよ"))
}