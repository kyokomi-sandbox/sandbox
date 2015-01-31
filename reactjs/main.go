package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"net/http"
	"github.com/zenazn/goji/web"
	"encoding/json"
	"io/ioutil"
	"io"
	"strings"
)

type Comment struct {
	Author string `json:"author"`
	Text string `json:"text"`
}

var comments []Comment

func init() {
	comments = make([]Comment, 0)
	comments = append(comments, Comment{Author: "Pete Hunt", Text: "Hey there!"})
}

func main() {
	fmt.Println("hello world")

	goji.Get("/comments.json", GetComments)
	goji.Post("/comments.json", PostComments)
	goji.Get("/*", http.FileServer(http.Dir("public")))

	goji.Serve()
}

func GetComments(_ web.C, w http.ResponseWriter, _ *http.Request) {
	data ,err := json.Marshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(data))
}

func PostComments(_ web.C, w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment Comment
	params := strings.Split(string(reqData), "&")
	for _, param := range params {
		keyVal := strings.Split(param, "=")
		switch keyVal[0] {
		case "author":
			comment.Author = keyVal[1]
		case "text":
			comment.Text = keyVal[1]
		}
	}
	comments = append(comments, comment)

	data ,err := json.Marshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, string(data))
}
