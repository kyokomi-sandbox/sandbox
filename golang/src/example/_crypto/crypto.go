package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	cryptoSHA1Example()
	cryptoSHA2Example()
	cryptoMD5Example()
}

func cryptoSHA1Example() {
	h := sha1.New()
	io.WriteString(h, "foo")
	fmt.Printf("SHA1   %x\n", h.Sum(nil))
}

func cryptoSHA2Example() {
	h224 := sha256.New224()
	io.WriteString(h224, "foo")
	h256 := sha256.New()
	io.WriteString(h256, "foo")

	h384 := sha512.New384()
	io.WriteString(h384, "foo")
	h512 := sha512.New()
	io.WriteString(h512, "foo")

	fmt.Printf("SHA224 %x\n", h224.Sum(nil))
	fmt.Printf("SHA256 %x\n", h256.Sum(nil))
	fmt.Printf("SHA384 %x\n", h384.Sum(nil))
	fmt.Printf("SHA512 %x\n", h512.Sum(nil))
}

func cryptoMD5Example() {
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
