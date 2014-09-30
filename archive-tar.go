package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func tarWrite() {

	f, err := os.Create("test.tar")
	if err != nil {
		log.Fatal(err)
	}
	tw := tar.NewWriter(f)

	baseDir := "./_test/"

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		// ディレクトリは一旦対象外とする
		if file.IsDir() {
			continue
		}

		// ヘッダー書き込み
		hdr := &tar.Header{
			Name: file.Name(),
			Size: file.Size(),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}

		// body書き込み
		data, err := ioutil.ReadFile(baseDir + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write(data); err != nil {
			log.Fatal(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func tarRead() {
	fmt.Println("Hello Go Sandbox!")

	// tarファイルを読み込み
	file, err := ioutil.ReadFile("test.tar")
	if err != nil {
		log.Fatal(err)
	}

	// tarのreaderを作成
	r := bytes.NewReader(file)
	tr := tar.NewReader(r)

	for {
		// ファイルの終わりまで読み込み
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// ファイル名とサイズを出力
		fmt.Printf("fileName : %s (%d)\n", hdr.Name, hdr.Size)

		// .txtファイルのみ内容を標準出力する
		if !strings.HasSuffix(hdr.Name, ".txt") {
			continue
		}
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
