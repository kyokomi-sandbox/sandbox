package service

import (
	"fmt"
	"net/http"
)

var Account AccountService

type AccountService struct {
	Name string
}

var _ Service = (*AccountService)(nil)

func (a AccountService) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(a.Name)
	w.Write([]byte("hoge"))
	a.Name = "hoge"
}
