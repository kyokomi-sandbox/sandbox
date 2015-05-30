package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/paypal"
)

type authKey string

func NewContext(ctx context.Context) context.Context {
	return WithPayPal(ctx)
}

func WithPayPal(ctx context.Context) context.Context {
	clientID := os.Getenv("PAYPAL_CLIENTID")
	if clientID == "" {
		// TODO: error?
		clientID = "AZz2Ai4gQliluOXFZCEt7ktV7m4cK9hilomvEzWq8Ojq6mKbT1WVFr9_Xv9XeXwq22sJzwgybttA2gyI"
	}
	secret := os.Getenv("PAYPAL_SECRET")
	if secret == "" {
		// TODO: error?
	}
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000/"
	}

	callBackURL := baseURL + "auth/paypal/callback"

	fmt.Println(callBackURL)

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: secret,
		RedirectURL:  callBackURL,
		Endpoint:     paypal.Endpoint,
		Scopes:       []string{"profile email address phone"},
	}

	// TODO: 本番と開発で呼び分け
	conf.Endpoint = paypal.SandboxEndpoint

	return context.WithValue(ctx, authKey("paypal"), conf)
}

func GetAuthToken(ctx context.Context, code string) (*oauth2.Token, error) {
	c, _ := FromContext(ctx)
	return c.Exchange(oauth2.NoContext, code)
}

func FromContext(ctx context.Context) (*oauth2.Config, bool) {
	conf, ok := ctx.Value(authKey("paypal")).(*oauth2.Config)
	return conf, ok
}
