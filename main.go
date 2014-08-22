package main

import "fmt"

func main() {
	fmt.Println("Hello Go Sandbox!")

	// ----------------------------------------------------------------
	// template
	text := CreateTemplateTree("Hello", "golang").Execute()
	fmt.Println("Result:", text)
	// ----------------------------------------------------------------
	// strings
	Sample()
	// ----------------------------------------------------------------
	// cli-scan
	c := CliScan {
		Scans: []Scan{
			{
				Key: "answer",
				Message: "Do you want to create one? [Y/n]",
			},
		},
	}
	fmt.Println("answer >", c.scan("answer"))
	// ----------------------------------------------------------------
	// Excel
	readExcel()
}

