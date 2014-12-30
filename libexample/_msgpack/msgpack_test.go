package main

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/ugorji/go/codec"
	msgpack "gopkg.in/vmihailenco/msgpack.v1"
)

/**
$ go test -bench .
PASS
BenchmarkMsgPackSampleMarshalMs  5000000               279 ns/op              64 B/op          1 allocs/op
BenchmarkMsgPackSampleAppendMsg 10000000               118 ns/op         177.59 MB/s           0 B/op          0 allocs/op
BenchmarkMsgPackSampleUnmarshal 10000000               204 ns/op         102.60 MB/s           0 B/op          0 allocs/op
BenchmarkMsgPackSampleEncode    10000000               136 ns/op         153.78 MB/s           0 B/op          0 allocs/op
BenchmarkMsgPackSampleDecode     3000000               422 ns/op          49.69 MB/s          16 B/op          2 allocs/op
BenchmarkMsgPackEncode______     2000000               800 ns/op          54.99 MB/s          48 B/op          1 allocs/op
BenchmarkMsgPackDecode______     1000000              1333 ns/op          33.00 MB/s          96 B/op          6 allocs/op
BenchmarkCodecEncode______       1000000              1360 ns/op          32.33 MB/s         128 B/op          2 allocs/op
BenchmarkCodecDecode______       1000000              1512 ns/op          29.10 MB/s          96 B/op         10 allocs/op
BenchmarkJsonEncode______        1000000              1123 ns/op          50.73 MB/s           8 B/op          1 allocs/op
BenchmarkJsonDecode______         500000              3111 ns/op          18.00 MB/s          32 B/op          3 allocs/op
ok      github.com/kyokomi-sandbox/go-sandbox/lib-example/msgpack       17.825s

*/

var m = MsgPackSample{
	Name:    "麻婆豆腐",
	Num:     80,
	Message: "AngelBeats!",
}

func BenchmarkMsgPackEncode______(b *testing.B) {
	var buf = &bytes.Buffer{}
	var enc = msgpack.NewEncoder(buf)

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
	var dec = msgpack.NewDecoder(buf)
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
