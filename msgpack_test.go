package main

import (
	"testing"
	"encoding/json"
	"bytes"
	msgpack "gopkg.in/vmihailenco/msgpack.v1"
	"log"
)

/**
BenchmarkCodec	  500000	      2961 ns/op
BenchmarkMsgPa	 1000000	      2218 ns/op
BenchmarkJsonc	  300000	      4725 ns/op
ok  	github.com/kyokomi/GoSandbox	5.240s
 */

var m = MsgPackSample{
	Name: "麻婆豆腐",
	Num: 80,
	Message: "AngelBeats!",
}

func BenchmarkMsgPackEncode(b *testing.B) {
	var buf   = &bytes.Buffer{}
	var enc   = msgpack.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(&m); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkMsgPackDecode(b *testing.B) {
	data, err := msgpack.Marshal(&m)
	if err != nil {
		log.Fatalln(err)
	}

	var buf = &bytes.Buffer{}
	var dec  = msgpack.NewDecoder(buf)
	var m2 MsgPackSample

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.Write(data)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkJsonEncode(b *testing.B) {

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(&m); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkJsonDecode(b *testing.B) {

	buf := &bytes.Buffer{}
	dec := json.NewDecoder(buf)
	var m2 MsgPackSample

	data, err := json.Marshal(&m)
	if err != nil {
		log.Fatalln(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.Write(data)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}
