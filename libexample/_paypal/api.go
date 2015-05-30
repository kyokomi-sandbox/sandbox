package main

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"github.com/k0kubun/pp"
)

type funcType string

type AuthCallbackFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, token *oauth2.Token)

func FromAuthCallbackFunc(ctx context.Context) (AuthCallbackFunc, bool) {
	conf, ok := ctx.Value(funcType("authCallbackFunc")).(AuthCallbackFunc)
	return conf, ok
}

func WithAuthCallbackFunc(ctx context.Context, authCallbackFunc AuthCallbackFunc) context.Context {
	return context.WithValue(ctx, funcType("authCallbackFunc"), authCallbackFunc)
}

func LoginPayPal(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c, _ := FromContext(ctx)
	http.Redirect(w, r, c.AuthCodeURL(""), 302)
}

func AuthPayPalCallback(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pp.Println(r)
	code := r.FormValue("code")
	token, err := GetAuthToken(ctx, code)
	if err != nil {
		log.Println("[ERROR] paypal.GetAuthToken", err)
		w.Write([]byte(err.Error()))
		return
	}

	fn, _ := FromAuthCallbackFunc(ctx)
	fn(ctx, w, r, token)
}
