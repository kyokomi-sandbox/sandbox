package main

import "fmt"

func main() {
	fmt.Println("Hello Go Sandbox!")

	text := CreateTemplateTree("Hello", "golang").Execute()

	fmt.Println("Result:", text)
}

