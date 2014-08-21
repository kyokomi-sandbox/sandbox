package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

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
	xlFile, err := xlsx.OpenFile("_test/sample2.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		// sheet単位でfile生成
		fmt.Println(sheet.Name)
		f, err := os.Create("_test/" + sheet.Name + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		// rowはまとめて1行にする
		for _, row := range sheet.Rows {
			fmt.Println(row.Cells)
			for idx, cell := range row.Cells {
				if idx != 0 {
					f.WriteString(",")
				}
				f.WriteString(cell.String())
			}
			f.WriteString("\n")
		}
		f.Close()
	}
}

