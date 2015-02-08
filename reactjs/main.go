package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"log"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/url"
)

type Comment struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

var fileMutex sync.RWMutex
var comments []Comment

func init() {

	// _comments.jsonを読み込む。もし1レコードもなかったら初期コメントを入れる

	data, err := ioutil.ReadFile("_comments.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(data, &comments); err != nil {
		log.Fatalln(err)
	}

	if len(comments) < 1 {
		comments = append(comments, Comment{Author: "Pete Hunt", Text: "Hey there!"})
	}
}

func main() {
	fmt.Println("hello world")

	goji.Get("/comments.json", GetComments)
	goji.Post("/comments.json", PostComments)
	goji.Get("/*", http.FileServer(http.Dir("public")))

	goji.Serve()
}

func GetComments(_ web.C, w http.ResponseWriter, _ *http.Request) {
	responseWriteComments(w)
}

func PostComments(_ web.C, w http.ResponseWriter, r *http.Request) {

	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Commentを追加
	comment := parseComment(reqData)
	comments = append(comments, comment)

	responseWriteComments(w)
}

func responseWriteComments(w http.ResponseWriter) {
	data, err := json.Marshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 雑にFileを更新（DBとかRedisでもいい）
	go func(data []byte) {
		fileMutex.Lock()
		defer fileMutex.Unlock()

		if err := ioutil.WriteFile("_comments.json", data, 0644); err != nil {
			log.Println("file write error ", err)
		}
	}(data)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(data))
}

// author=foo&text=bar の形式でpostされる
func parseComment(data []byte) Comment {
	var comment Comment
	params := strings.Split(string(data), "&")
	for _, param := range params {
		keyVal := strings.Split(param, "=")
		switch keyVal[0] {
		case "author":
			comment.Author = keyVal[1]
		case "text":
			comment.Text, _ = url.QueryUnescape(keyVal[1])
		}

		fmt.Println("comment: ", string(keyVal[1]))
	}
	return comment
}
