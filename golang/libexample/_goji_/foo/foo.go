package foo

import (
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goji_/context"
	"github.com/zenazn/goji/web"
	"net/http"
	"fmt"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_goji_/service"
)


type FooService struct {
	context *context.Context
}
var _ service.Service = (*FooService)(nil)

func NewService(c *context.Context) *FooService {
	f := FooService{}
	f.context = c

	return &f
}

func (f FooService) HelloHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s FooHello, %s!", f.context.Name(), c.URLParams["name"])
}
