package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	readExcel("./sample/sample2.xlsx")
}

func readExcel(readFilePath string) {
	xlFile, err := xlsx.OpenFile(readFilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		// sheet単位でfile生成
		fmt.Println(sheet.Name)
		f, err := os.Create("./sample/" + sheet.Name + ".txt")
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
