package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println("Hello Go Sandbox!")

	var buf bytes.Buffer

	buf.WriteString("hogehoge")
	fmt.Println(string(buf.Bytes()))

	bufTitle := bytes.ToTitle(buf.Bytes())
	fmt.Println(string(bufTitle))
}
