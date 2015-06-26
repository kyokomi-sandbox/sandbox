package main

import (
	"net/http"
	"github.com/kyokomi-sandbox/go-sandbox/other/waf/service"
)

func main() {
	http.HandleFunc("/", service.Account.Index)

	http.ListenAndServe(":8000", nil)
}
