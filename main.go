package main

import "fmt"

type CliScan struct {
	Scans []Scan
}
type Scan struct {
	Key string
	Message string
	Value string
}

func (c *CliScan) scan(key string) string {
	var s  Scan
	for _, s = range c.Scans {
		if s.Key == key {
			break
		}
	}

	if s.Key != key {
		return ""
	}

	fmt.Print(s.Message, ": ")
	fmt.Scanln(&s.Value)
	return s.Value
}

func main() {
	fmt.Println("Hello Go Sandbox!")

	c := CliScan {
		Scans: []Scan{
			{
				Key: "answer",
				Message: "Do you want to create one? [Y/n]",
			},
		},
	}

//	// template
//	text := CreateTemplateTree("Hello", "golang").Execute()
//	fmt.Println("Result:", text)
//
//	// strings
//	Sample()

//	cli.NewApp()
//	fmt.Println("You don't have any configuration file")
//	fmt.Print("Do you want to create one? [Y/n]: ")
//	var answer string
//	fmt.Scanln(&answer)
//
//	fmt.Println("answer >", answer)
	fmt.Println("answer >", c.scan("answer"))
}

