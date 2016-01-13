package main

import (
	"net/http"

	"github.com/zenazn/goji/web"
	"golang.org/x/oauth2"
)

var conf *oauth2.Config

func initFacebook(appID, secret string) {
	conf = &oauth2.Config{
		ClientID:     appID,
		ClientSecret: secret,
		RedirectURL:  "http://localhost:8000/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
		Scopes: []string{"email"},
	}
}

func FacebookOAuthLogin(ctx *web.C, handler http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx.Env["facebook"] = conf
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(h)
}

func GetFaceBook(ctx *web.C) *oauth2.Config {
	c, ok := ctx.Env["facebook"].(*oauth2.Config)
	if !ok {
		panic("not facebook config")
	}
	return c
}
