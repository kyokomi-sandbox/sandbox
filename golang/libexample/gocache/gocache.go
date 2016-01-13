package main

import (
	"time"

	"fmt"

	"github.com/patrickmn/go-cache"
)

type Hoge struct {
	Name string
	Num  int
}

func main() {
	c := cache.New(5*time.Minute, 30*time.Second)

	c.Set("foo", "bar", cache.DefaultExpiration)

	fmt.Println(c.Get("foo"))

	h := &Hoge{
		Name: "fuga",
		Num:  99999,
	}
	c.Set("foo", h, cache.DefaultExpiration)

	// Setした時点で別にPointerになるらしい
	if ch, ok := c.Get("foo"); ok {
		x := ch.(*Hoge)
		fmt.Printf("%v %v\n", &h, &x)
	}
}
