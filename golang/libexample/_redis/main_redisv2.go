package main

import (
	redis "gopkg.in/redis.v2"
	"fmt"
	"flag"
	"strings"
)

func main() {

    example2()
}

func example1() {

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

func example2() {

    var host string
    flag.StringVar(&host, "host", "localhost", "host name")
    flag.Parse()

    opt := redis.Options{}
    opt.Addr = strings.Join([]string{host, "6379"}, ":")
    opt.Network = "tcp"

    client := redis.NewClient(&opt)
    defer func() {
        fmt.Println("defer")
    }()

    client.Set("foo12", "bar")
    strCmd := client.Get("foo1")
    fmt.Println(strCmd.String())
    
    panic("error")

    client.Close()
}