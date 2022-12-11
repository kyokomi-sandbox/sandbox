package main

import (
    "context"
    "log"
    "net/http"

    greetv1 "example/gen/greet/v1"
    "example/gen/greet/v1/greetv1connect"

    "github.com/bufbuild/connect-go"
)

func main() {
    client := greetv1connect.NewGreetServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
				// By default, connect-go servers support ingress from all three protocols without any configuration.
				// connect-go clients use the Connect protocol by default, 
				// but can use either the gRPC or gRPC-Web protocols by setting the WithGRPC or WithGRPCWeb client options.
				// connect.WithGRPC(), 
    )
    res, err := client.Greet(
        context.Background(),
        connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println(res.Msg.Greeting)
}

