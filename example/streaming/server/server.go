package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foreverYoung)
	http.ListenAndServe(":8000", nil)
}

func foreverYoung(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)

	// TODO: 5秒繰り返す適当なコード
	for i := 0; i < 5; i++ {
		w.Write([]byte("young\r\n"))
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, "hello end")
}
