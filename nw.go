package main

import (
	"fmt"
	nw "github.com/lonnc/golang-nw"
	"net/http"
	"log"
)

/*
golang-nw

https://github.com/lonnc/golang-nw

nodewebkitの中身をgolangで書くやつ。

# 導入手順

1. nodewebkit自体のinstall
`$ npm install -g nodewebkit`

2. golang-nw-pkgのinstall
`$ go install github.com/lonnc/golang-nw/cmd/golang-nw-pkg`

3. golang-nw-pkgで動かすgolangのinstall
`$ go install github.com/lonnc/golang-nw/cmd/example`

4. golang-nw-pkgでbuildする
-appにinstallしたgoのバイナリのパスを指定する
`golang-nw-pkg -app=/Users/kyokomi/bin/example -name="My Application" -bin="myapp2" -toolbar=false`

# メモ
- install手順がちょっとわかりにくかった
- git cloneした時のデフォルトnode-webkitのバージョンがDownloadできなくて「？？？」ってなった
- golang-nw-buildで組み込んでる`script.js`でエラーがでてる
	- 次回起動時用に画面サイズの保持でlocalStrageを使ってるところでエラーってるみたい？

```
[11870:1229/112845:INFO:CONSOLE(9)] "Uncaught TypeError: Cannot read property 'width' of null", source: file:///var/folders/_9/7ppk1m3934db1h84qrxlvf980000gn/T/.org.chromium.Chromium.BdrNIn/script.js (9)
[11870:1229/112848:INFO:CONSOLE(42)] "Uncaught TypeError: Cannot set property 'x' of null", source: file:///var/folders/_9/7ppk1m3934db1h84qrxlvf980000gn/T/.org.chromium.Chromium.BdrNIn/script.js (42)
```

 */
func nwExample() {

	// Setup our handler
	http.HandleFunc("/", hello2)

	// Create a link back to node-webkit using the environment variable
	// populated by golang-nw's node-webkit code
	nodeWebkit, err := nw.New()
	if err != nil {
		panic(err)
	}


	// Pick a random localhost port, start listening for http requests using default handler
	// and send a message back to node-webkit to redirect
	if err := nodeWebkit.ListenAndServe(nil); err != nil {
		log.Fatalln(err)
	}
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from golang.")
}
