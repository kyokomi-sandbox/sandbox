package main

import (
	"fmt"

	redis "gopkg.in/redis.v2"
)

var redisDB *redis.Client

func initRedis() {
	redisDB = redis.NewTCPClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "localhost", 6379),
		DB:   int64(1), // use default DB
	})

	fmt.Println(redisDB)
}
