package main

import (
	"os"
	"errors"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/paypal"
)

type authKey string

var (
	ErrNonClientID = errors.New("env error PAYPAL_CLIENTID")
	ErrNonSecret   = errors.New("env error PAYPAL_SECRET")
	ErrNonBaseURL  = errors.New("env error BASE_URL")
)

func NewContext(ctx context.Context) context.Context {
	ctx, err := WithPayPal(ctx)
	if err != nil {
		log.Println(err)
	}
	return ctx
}

func WithPayPal(ctx context.Context) (context.Context, error) {
	clientID := os.Getenv("PAYPAL_CLIENTID")
	if clientID == "" {
		return ctx, ErrNonClientID
	}
	secret := os.Getenv("PAYPAL_SECRET")
	if secret == "" {
		return ctx, ErrNonSecret
	}
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		return ctx, ErrNonBaseURL
	}

	// TODO: callbackを外から指定
	callBackURL := baseURL + "auth/paypal/callback"

	// TODO: scopeは外から指定できるようにする
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: secret,
		RedirectURL:  callBackURL,
		Endpoint:     paypal.Endpoint,
		Scopes: []string{
			"profile",
			"email",
			"address",
			"phone",
			"https://uri.paypal.com/services/paypalattributes",
		},
	}

	// TODO: 本番と開発で呼び分ける
	conf.Endpoint = paypal.SandboxEndpoint

	return context.WithValue(ctx, authKey("paypal"), conf), nil
}

func GetAuthToken(ctx context.Context, code string) (*oauth2.Token, error) {
	c, _ := FromContext(ctx)
	return c.Exchange(oauth2.NoContext, code)
}

func FromContext(ctx context.Context) (*oauth2.Config, bool) {
	conf, ok := ctx.Value(authKey("paypal")).(*oauth2.Config)
	return conf, ok
}
