package main

import "fmt"

func main() {
	c := CliScan{
		Scans: []Scan{
			Scan{Key: "hoge", Message: "test", Value: "foo"},
		},
	}
	fmt.Println("Scan: ", c.scan("hoge"))
}
