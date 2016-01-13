package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type funcType string

const (
	authCallbackFuncType funcType = "authCallbackFunc"
	authErrorFuncType    funcType = "authErrorFunc"
)

type AuthCallbackFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, token *oauth2.Token)
type AuthErrorFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)

func FromAuthCallbackFunc(ctx context.Context) (AuthCallbackFunc, bool) {
	conf, ok := ctx.Value(authCallbackFuncType).(AuthCallbackFunc)
	return conf, ok
}

func WithAuthCallbackFunc(ctx context.Context, authCallbackFunc AuthCallbackFunc) context.Context {
	return context.WithValue(ctx, authCallbackFuncType, authCallbackFunc)
}

func WithAuthErrorFunc(ctx context.Context, authCallbackFunc AuthErrorFunc) context.Context {
	return context.WithValue(ctx, authErrorFuncType, authCallbackFunc)
}

func FromAuthErrorFunc(ctx context.Context) (AuthErrorFunc, bool) {
	conf, ok := ctx.Value(authErrorFuncType).(AuthErrorFunc)
	return conf, ok
}

func LoginPayPal(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c, _ := FromContext(ctx)
	http.Redirect(w, r, c.AuthCodeURL(""), 302)
}

func AuthPayPalCallback(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := GetAuthToken(ctx, code)
	if err != nil {
		log.Println("[ERROR] paypal.GetAuthToken", err)
		if fn, ok := FromAuthErrorFunc(ctx); ok {
			fn(ctx, w, r, err)
		}
		return
	}

	if fn, ok := FromAuthCallbackFunc(ctx); ok {
		fn(ctx, w, r, token)
	}
}
