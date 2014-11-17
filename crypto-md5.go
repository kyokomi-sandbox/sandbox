package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func sampleMD5() {
	h := md5.New()
	io.WriteString(h, "hello world")
	result := h.Sum(nil)

	fmt.Printf("%x\n", result)

	if fmt.Sprintf("%x", result) == "5eb63bbbe01eeed093cb22bb8f5acdc3" {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
