package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile)

	f, err := os.OpenFile("hoge.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("hoge"); err != nil {
		panic(err)
	}
}
