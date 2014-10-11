package main

import (
	"fmt"
	"bytes"
)

func bytesExec() {
	var buf bytes.Buffer

	buf.WriteString("hogehoge")
	fmt.Println(string(buf.Bytes()))

	bufTitle := bytes.ToTitle(buf.Bytes())
	fmt.Println(string(bufTitle))
}
