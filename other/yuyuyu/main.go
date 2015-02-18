package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"os"
)

type Settings struct {
	RootURL string
	APIPath string
}

var settings = Settings{
	RootURL: "http://www.onsen.ag/",
	APIPath: "http://www.onsenradio.info/data/api/getMovieInfo/yuyuyu",
}

type yuyuyu struct {
	AllowExpand string        `json:"allowExpand"`
	Cm          []interface{} `json:"cm"`
	Copyright   string        `json:"copyright"`
	Count       string        `json:"count"`
	Guest       string        `json:"guest"`
	Link        []struct {
		ImagePath string `json:"imagePath"`
		URL       string `json:"url"`
	} `json:"link"`
	Mail      string `json:"mail"`
	MoviePath struct {
		Android string `json:"Android"`
		IPhone  string `json:"iPhone"`
		PC      string `json:"pc"`
	} `json:"moviePath"`
	OptionText     string `json:"optionText"`
	Personality    string `json:"personality"`
	RecommendGoods []struct {
		ImagePath string `json:"imagePath"`
		URL       string `json:"url"`
	} `json:"recommendGoods"`
	RecommendMovie []interface{} `json:"recommendMovie"`
	Schedule       string        `json:"schedule"`
	ThumbnailPath  string        `json:"thumbnailPath"`
	Title          string        `json:"title"`
	Type           string        `json:"type"`
	Update         string        `json:"update"`
	URL            string        `json:"url"`
}

func main() {

	yuyu, err := metadata()
	if err != nil {
		log.Fatalln(err)
	}

	if err := downloadMovie(yuyu.MoviePath.PC, yuyu.URL, yuyu.Count+".mp3"); err != nil {
		log.Fatalln(err)
	}
}

func metadata() (*yuyuyu, error) {
	res, err := http.Get(settings.APIPath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	callback := string(data)

	callback = strings.Replace(callback, "callback(", "", 1)
	callback = strings.Replace(callback, ");", "", 1)

	var yuyu yuyuyu
	if err := json.Unmarshal(bytes.NewBufferString(callback).Bytes(), &yuyu); err != nil {
		return nil, err
	}

	return &yuyu, nil
}

func downloadMovie(movieURL string, dirPath, fileName string) error {

	res, err := http.Get(movieURL)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Println(err)
	}

	f, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		return err
	}

	f.Write(data)

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
