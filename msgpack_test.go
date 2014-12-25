package main

import (
	"testing"
	"encoding/json"
	"bytes"
)

/**
BenchmarkCodec	  500000	      2961 ns/op
BenchmarkMsgPa	 1000000	      2218 ns/op
BenchmarkJsonc	  300000	      4725 ns/op
ok  	github.com/kyokomi/GoSandbox	5.240s
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

func BenchmarkJsonc(b *testing.B) {

	buf := &bytes.Buffer{}
	bufR := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(bufR)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(&m)

		bufR.Write(buf.Bytes())

		var m2 MsgPackSample
		dec.Decode(&m2)

		bufR.Reset()
	}
}


