package main

import (
	"log"
	"net/http"

	"github.com/gunosy/kami"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"golang.org/x/net/context"
)

func main() {
	secretKey := "kyokomi-secret"

	gomniauth.SetSecurityKey(secretKey)
	gomniauth.WithProviders(
		github.New(githubClientID, githubSecretKey, githubRedirectURL),
		google.New(googleClientID, googleSecretKey, googleRedirectURL),
	)

	kami.Get("/", indexHandler)
	kami.Get("/auth/login/:provider", loginHandler)
	kami.Get("/auth/callback/:provider", callbackHandler)

	kami.Serve()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func loginHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	providerName := kami.Param(ctx, "provider")

	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		log.Fatalln("認証プロバイダー の取得に失敗しました:", provider, "-", err)
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("GetBeginAuthURLの呼び出し中にエラーが発生しました:", provider, "-", err)
	}
	w.Header().Set("Location", loginUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func callbackHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	providerName := kami.Param(ctx, "provider")

	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		log.Fatalln("認証プロバイダーの取得に失敗しました", provider, "-", err)
	}
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	if err != nil {
		log.Fatalln("認証を完了できませんでした", provider, "-", err)
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		log.Fatalln("ユーザーの取得に失敗しました", provider, "- ", err)
	}

	authCookieValue := objx.New(map[string]interface{}{
		"name": user.Name(),
	}).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/"})
	log.Println("######### ", authCookieValue)

	w.Header()["Location"] = []string{"/"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}
