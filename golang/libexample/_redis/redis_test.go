package main

import (
	"testing"
	redistest "github.com/soh335/go-test-redisserver"
	redis "gopkg.in/redis.v2"
	"fmt"
)

func TestMain(t *testing.T) {

	server ,err := redistest.NewServer(true, nil)
	if err != nil {
		t.Error(err)
	}
	defer server.Stop()

	host := server.Config["unixsocket"]

	opt := redis.Options{}
	opt.Addr = host
	opt.Network = "unix"

	client := redis.NewClient(&opt)
	defer client.Close()

	if strCmd := client.Set("foo2", "bar"); strCmd.Err() != nil {
		t.Errorf(strCmd.Err())
	}

	if strCmd := client.Get("foo2"); strCmd.Err() != nil {
		t.Errorf(strCmd.Err())
	} else {
		fmt.Println(strCmd.String())
	}
}
