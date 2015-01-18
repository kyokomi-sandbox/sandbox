package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	tabWriterSample()
}

func tabWriterSample() {
	// 4タブ
	w1 := tabwriter.NewWriter(os.Stdout, 0, 4, 0, '\t', 0)
	io.WriteString(w1, "a\tbcccs\tc\td\te\n")
	w1.Flush()

	// 4スペースタブ
	w2 := tabwriter.NewWriter(os.Stdout, 4, 0, 0, ' ', 0)
	io.WriteString(w2, "a\tbcccs\tc\td\te\n")
	w2.Flush()

	// 8スペースタブ右寄せ
	w3 := tabwriter.NewWriter(os.Stdout, 8, 0, 0, ' ', tabwriter.AlignRight)
	io.WriteString(w3, "a\tbcccs\tc\td\te\n")
	w3.Flush()

	// 5スペースタブ
	w4 := tabwriter.NewWriter(os.Stdout, 5, 0, 0, ' ', 0)
	io.WriteString(w4, "aaaaa\tbbbbb\tccccc\tddddd\teeeee\n")
	w4.Flush()

	// 5スペースタブ1パディング
	w5 := tabwriter.NewWriter(os.Stdout, 5, 0, 1, ' ', 0)
	io.WriteString(w5, "aaaaa\tbbbbb\tccccc\tddddd\teeeee\n")
	w5.Flush()

	fmt.Println(CreateTemplateTree("hoge", "fuga").Execute())
}
