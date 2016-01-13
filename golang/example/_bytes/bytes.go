package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("hogehoge")
	fmt.Println(string(buf.Bytes()))

	bufTitle := bytes.ToTitle(buf.Bytes())
	fmt.Println(string(bufTitle))
}
