package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
)

func main() {

	res, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
	}
}
