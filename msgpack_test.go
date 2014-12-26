package main

import (
	"testing"
	"encoding/json"
	"bytes"
	msgpack "gopkg.in/vmihailenco/msgpack.v1"
	"log"
	"github.com/ugorji/go/codec"
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

func BenchmarkMsgPackEncode______(b *testing.B) {
	var buf   = &bytes.Buffer{}
	var enc   = msgpack.NewEncoder(buf)

	if err := enc.Encode(&m); err != nil {
		log.Fatalln(err)
	}
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(&m); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkMsgPackDecode______(b *testing.B) {
	data, err := msgpack.Marshal(&m)
	if err != nil {
		log.Fatalln(err)
	}

	var buf = &bytes.Buffer{}
	var dec  = msgpack.NewDecoder(buf)
	var m2 MsgPackSample

	buf.Write(data)
	b.SetBytes(int64(buf.Len()))
	
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.Write(data)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkCodecEncode______(b *testing.B) {
	buf := &bytes.Buffer{}
	var mh = &codec.MsgpackHandle{RawToString: true}
	var enc = codec.NewEncoder(buf, mh)

	if err := enc.Encode(&m); err != nil {
		log.Fatalln(err)
	}
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(&m); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkCodecDecode______(b *testing.B) {
	buf := &bytes.Buffer{}
	var mh = &codec.MsgpackHandle{RawToString: true}
	var dec = codec.NewDecoder(buf, mh)

	var m2 MsgPackSample
	buf2 := &bytes.Buffer{}
	enc := codec.NewEncoder(buf2, mh)
	enc.Encode(&m)

	buf.Write(buf2.Bytes())
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.Write(buf2.Bytes())

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkJsonEncode______(b *testing.B) {

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	if err := enc.Encode(&m); err != nil {
		log.Fatalln(err)
	}
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		if err := enc.Encode(&m); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkJsonDecode______(b *testing.B) {

	buf := &bytes.Buffer{}
	dec := json.NewDecoder(buf)
	var m2 MsgPackSample

	data, err := json.Marshal(&m)
	if err != nil {
		log.Fatalln(err)
	}

	buf.Write(data)
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.Write(data)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}
