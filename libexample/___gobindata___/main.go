package main

import (
	"log"
	"fmt"
)

//go:generate go-bindata data/

func main () {

	data, err := Asset("data/emoji.json")
	if err != nil {
		// Asset was not found.
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
