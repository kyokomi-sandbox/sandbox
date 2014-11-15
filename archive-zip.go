package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readerExample() {
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		// 30 byteだけcopy
		if _, err := io.CopyN(os.Stdout, rc, 30); err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func writerExample() {
	f, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	zw := zip.NewWriter(f)

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

		f, err := zw.Create(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		// body書き込み
		data, err := ioutil.ReadFile(baseDir + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(data); err != nil {
			log.Fatal(err)
		}
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
}
