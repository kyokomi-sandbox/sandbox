package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
)

type Settings struct {
	RootURL string
	APIPath string
}

var settings = Settings{
	RootURL: "http://www.onsen.ag/",
	APIPath: "http://www.onsenradio.info/data/api/getMovieInfo/yuyuyu",
}

func main() {

	res, err := http.Get(settings.APIPath)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	callback := string(data)

	callback = strings.Replace(callback, "callback(", "", 1)
	callback = strings.Replace(callback, ");", "", 1)

	var jsonRaw interface{}
	if err := json.Unmarshal(bytes.NewBufferString(callback).Bytes(), &jsonRaw); err != nil {
		log.Fatalln(err)
	}

	pp.Println(jsonRaw)
}
