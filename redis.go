package main

import (
	"github.com/hoisie/redis"
	"fmt"
	"time"
	"log"
)

const ONE_WEEK_DAY = 7
const VOTE_SCORE = 432

func RedisExample(article string) {
	var client redis.Client

	cutoff := time.Now().Add(ONE_WEEK_DAY * time.Second)

	hit, err := client.Exists("time:")
	if err != nil {
		log.Fatal("Existsエラーだよ", err.Error())
	}

	if !hit {
		log.Fatal("記事がないよ ")
	}

	score, err := client.Zscore("time:", []byte(article))
	if err != nil {
		log.Fatal("Zscoreエラーだよ", err.Error())
	}

	if cutoff.After(time.Unix(int64(score), 0)) {
		fmt.Println("Afterだった")
		return
	}

	var key = "hello"

	client.Set(key, []byte("world"))
	val, _ := client.Get("hello")
	fmt.Println(key, string(val))
}

