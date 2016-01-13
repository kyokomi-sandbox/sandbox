package main

import (
	"net/http"

	"log"

	"fmt"

	"github.com/k0kubun/pp"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"golang.org/x/oauth2"
	redis "gopkg.in/redis.v2"
	"flag"
	"github.com/go-xweb/uuid"
)

// 参考: http://umegusa.hatenablog.jp/entry/2015/03/09/002524

func main() {
	var appID, secret string
	flag.StringVar(&appID, "appid", "", "")
	flag.StringVar(&secret, "secret", "", "")
	flag.Parse()
	fmt.Println(appID, secret)

	initFacebook(appID, secret)
	initRedis()

	goji.Use(FacebookOAuthLogin)

	goji.Get("/", IndexHandler)
	goji.Get("/login", LoginHandler)
	goji.Get("/auth/callback", AuthCallbackHandler)
	goji.Serve()
	defer redisDB.Close()
}

func IndexHandler(ctx web.C, w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	if userID == "" {
		userID = r.Header.Get("MEET_APP_USER_ID")
	}
	token, err := redisDB.Get("auth:" + userID).Result()
	if err == redis.Nil {
		fmt.Fprint(w, "not login")
		return
	}

	if err != nil {
		log.Fatalln(err)
	}

	pp.Println(r.Header)
	pp.Println(r.Cookies())

	fmt.Println(token)
}

func LoginHandler(ctx web.C, w http.ResponseWriter, r *http.Request) {
	c := GetFaceBook(&ctx)

	url := c.AuthCodeURL("")
	http.Redirect(w, r, url, 301)
}

func AuthCallbackHandler(ctx web.C, w http.ResponseWriter, r *http.Request) {
	c := GetFaceBook(&ctx)

	code := r.FormValue("code")
	token, err := c.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(token.AccessToken, token.Expiry)

	// TODO: メアドとUUIDを紐付ける
	userID := uuid.New()
	_, err = redisDB.Set("auth:"+userID, token.AccessToken).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, userID)
}
