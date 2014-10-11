package main

import (
	"fmt"
	"compress/bzip2"
	"os"
	"log"
	"io/ioutil"
)

func bzip2Exec() {
	bz2File, err := os.Open("text.tar.bz2")
	if err != nil {
		log.Fatal(err)
	}
	defer bz2File.Close()

	f := bzip2.NewReader(bz2File)

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
