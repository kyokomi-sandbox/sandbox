package main

import (
	"fmt"
	"archive/tar"
	"io/ioutil"
	"log"
	"os"
)


func main() {
	fmt.Println("Hello Go Sandbox!")

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

