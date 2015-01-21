package sample

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

/*
$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkMegaJsonEncode	 1000000	      1430 ns/op	 247.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkJsonEncode____	  500000	      2778 ns/op	 127.79 MB/s	       8 B/op	       1 allocs/op
BenchmarkMegaJsonDecode	  200000	      7503 ns/op	  58.24 MB/s	     384 B/op	      33 allocs/op
BenchmarkJsonDecode____	  100000	     13083 ns/op	  33.40 MB/s	     293 B/op	       8 allocs/op
*/

var sampleJson = `
{
        "id": 556621823429722100,
        "id_str": "556621823429722112",
        "media_url": "http://pbs.twimg.com/media/B7mEmWvCQAA10cT.jpg",
        "media_url_https": "https://pbs.twimg.com/media/B7mEmWvCQAA10cT.jpg",
        "url": "http://t.co/ywJYwZQbv7",
        "display_url": "pic.twitter.com/ywJYwZQbv7",
        "expanded_url": "http://twitter.com/kyokomidev/status/556621824109211649/photo/1",
        "type": "photo"
}
`
func BenchmarkMegaJsonEncode(b *testing.B) {

	var m Media
	if err := json.Unmarshal([]byte(sampleJson), &m); err != nil {
		log.Fatalln(err)
	}

	buf := &bytes.Buffer{}
	enc := NewMediaJSONEncoder(buf)

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


func BenchmarkJsonEncode____(b *testing.B) {

	var m Media
	if err := json.Unmarshal([]byte(sampleJson), &m); err != nil {
		log.Fatalln(err)
	}

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

func BenchmarkMegaJsonDecode(b *testing.B) {

	buf := &bytes.Buffer{}
	dec := NewMediaJSONDecoder(buf)
	var m2 *Media

	buf.WriteString(sampleJson)
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.WriteString(sampleJson)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}

func BenchmarkJsonDecode____(b *testing.B) {

	buf := &bytes.Buffer{}
	dec := json.NewDecoder(buf)
	var m2 Media

	buf.WriteString(sampleJson)
	b.SetBytes(int64(buf.Len()))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buf.WriteString(sampleJson)

		if err := dec.Decode(&m2); err != nil {
			log.Fatalln(err)
		}
	}
}
