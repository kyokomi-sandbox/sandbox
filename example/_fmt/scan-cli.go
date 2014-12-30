package main

import "fmt"

type CliScan struct {
	Scans []Scan
}
type Scan struct {
	Key     string
	Message string
	Value   string
}

func (c *CliScan) scan(key string) string {
	var s Scan
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
