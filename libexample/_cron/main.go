package main

import (
	"fmt"

	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("5 * * * * *", func() { fmt.Println("Every hour on the half hour") }) // 毎分5秒
	c.AddFunc("*/5 * * * * *", func() { fmt.Println("Every hour on the half hour") }) // 5秒毎
	c.Start()
	defer c.Stop()

	time.Sleep(1 * time.Minute)
}
