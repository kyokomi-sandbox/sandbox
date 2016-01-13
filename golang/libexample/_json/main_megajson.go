package main

import (
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_json/sample"
	"bytes"
	"fmt"
)

//go:generate ./sample/sample.go

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

func main() {

	var reader bytes.Buffer
	reader.WriteString(sampleJson)

	var m *sample.Media
	sample.NewMediaJSONDecoder(&reader).Decode(&m)

	fmt.Println(m)
}
