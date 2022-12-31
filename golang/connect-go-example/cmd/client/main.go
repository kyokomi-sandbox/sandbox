package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"example/cmd/interceptor"
	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"

	"github.com/bufbuild/connect-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func extractRetryInfo(err error) (*errdetails.RetryInfo, bool) {
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		return nil, false
	}
	log.Println(connectErr.Details())

	for _, detail := range connectErr.Details() {
		msg, valueErr := detail.Value()
		if valueErr != nil {
			// Usually, errors here mean that we don't have the schema for this
			// Protobuf message.
			continue
		}

		if retryInfo, ok := msg.(*errdetails.RetryInfo); ok {
			return retryInfo, true
		}
	}
	return nil, false
}

func main() {
	interceptors := connect.WithInterceptors(interceptor.NewAuthInterceptor())

	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		// By default, connect-go servers support ingress from all three protocols without any configuration.
		// connect-go clients use the Connect protocol by default,
		// but can use either the gRPC or gRPC-Web protocols by setting the WithGRPC or WithGRPCWeb client options.
		//connect.WithGRPC(),
		interceptors,
	)
	ctx := context.Background()
	req := connect.NewRequest(&greetv1.GreetRequest{Name: "aaa"})
	req.Header().Add("Acme-Tenant-Id", "1234")

	res, err := client.Greet(ctx, req)
	if err != nil {
		log.Println(connect.CodeOf(err))
		log.Println(extractRetryInfo(err))
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
