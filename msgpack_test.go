package main

import (
	"testing"
)

/**
BenchmarkCodec	  500000	      2911 ns/op
BenchmarkMsgPa	 1000000	      2269 ns/op
ok  	github.com/kyokomi/GoSandbox	3.798s
 */

var m = MsgPackSample{
	Name: "ユイにゃん",
	Num: 17,
	Message: "AngelBeats!",
}

func BenchmarkCodec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		packUnPackCodec(m)
	}
}

func BenchmarkMsgPa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		packUnPackMsgPack(m)
	}
}




