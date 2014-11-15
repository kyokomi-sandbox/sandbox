package main

import (
	"github.com/hoisie/redis"
	"fmt"
	"time"
	"log"
	"strings"
	"strconv"
)

const ONE_WEEK_DAY = 7
const VOTE_SCORE = 432

func RedisExample(user, article string) {
	var client redis.Client

	postArticle(client, user, "hoge", "http://google.com")
	articleVote(client, user, article)
}

func articleVote(client redis.Client, user, article string) {

	cutoff := time.Now().Add(ONE_WEEK_DAY * time.Second)

	// 記事チェック
	if hit, err := client.Exists("time:"); err != nil {
		log.Fatal("Existsエラーだよ", err.Error())
	} else if !hit {
		log.Fatal("記事がないよ ")
	}

	// 期限チェック
	if score, err := client.Zscore("time:", []byte(article)); err != nil {
		log.Fatal("Zscoreエラーだよ", err.Error())
	} else {
		if cutoff.After(time.Unix(int64(score), 0)) {
			fmt.Println("Afterだった ", cutoff.Unix(), " ", int64(score))
			return
		}
	}

	articleID := strings.TrimPrefix(article, "article:")
	fmt.Println("articleID", articleID)

	ok , err := client.Sadd("voted:" + articleID, []byte(user))
	if err != nil {
		log.Fatal("Saddエラーだよ", err.Error())
	}

	if ok {
		if _, err := client.Zincrby("score:", []byte(article), VOTE_SCORE); err != nil {
			log.Fatal("Zincrbyエラーだよ", err.Error())
		}
		if _, err := client.Hincrby(article, "votes", 1); err != nil {
			log.Fatal("Hincrbyエラーだよ", err.Error())
		}
	}
}

func postArticle(client redis.Client, user, title, link string) string {
	id, err := client.Incr("article:")
	if err != nil {
		log.Fatal("Incrエラーだよ", err.Error())
	}

	articleID := strconv.FormatInt(id, 20)
	voted := "voted:" + articleID

	if _, err := client.Sadd(voted, []byte(user)); err != nil {
		log.Fatal("Saddエラーだよ", err.Error())
	}

	if _, err := client.Expire(voted, int64(ONE_WEEK_DAY * time.Second)); err != nil {
		log.Fatal("Expireエラーだよ", err.Error())
	}
	now := time.Now()
	article := "article:" + articleID
	fmt.Println("articleID", articleID)

	if err := client.Hmset(article, map[string]string{
			"title": title,
			"link": link,
			"poster": user,
			"time": now.String(),
			"votes": "1",
		}); err != nil {
		log.Fatal("Hmsetエラーだよ", err.Error())
	}

	if _, err := client.Zadd("score:", []byte(article), float64(now.Unix() + VOTE_SCORE)); err != nil {
		log.Fatal("Zaddエラーだよ", err.Error())
	}
	if _, err := client.Zadd("time:", []byte(article), float64(now.Unix())); err != nil {
		log.Fatal("Zaddエラーだよ", err.Error())
	}

	return articleID
}

func getArticle(client redis.Client, page int, order string) []map[string]string {
	if order == "" {
		order = "score:"
	}

	// TODO: 未実装

	return []map[string]string{}
}

func addRemoveGroups(client redis.Client, articleID int, toAdd, toRemove []string) {

	// TODO: 未実装
}

func getGroupArticle(client redis.Client, group string, page int) []map[string]string {
	order := "score:"

	key := order + group

	// TODO: 未実装

	return getArticle(client, page, key)
}
