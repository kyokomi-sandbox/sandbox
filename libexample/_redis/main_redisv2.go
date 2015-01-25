package main

import (
	redis "gopkg.in/redis.v2"
	"fmt"
	"flag"
	"strings"
)

func main() {

	var host string
	flag.StringVar(&host, "host", "localhost", "host name")
	flag.Parse()

	opt := redis.Options{}
	opt.Addr = strings.Join([]string{host, "6379"}, ":")
	opt.Network = "tcp"

	client := redis.NewClient(&opt)
	defer client.Close()

	client.Set("foo", "bar")

	strCmd := client.Get("foo")
	fmt.Println(strCmd.String())
}
