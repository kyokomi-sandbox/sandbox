package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	scanLines(strings.NewReader("hogehogehoge\nfugafuga"))

	scanCustomSplit()

	bufioWriter()
}

func scanLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func scanCustomSplit() {
	const input = "1234	5678	123456790123	4567890"
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\t'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func bufioWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "hogehogehoge")

	time.Sleep(time.Second * 1)
	w.Flush()
}
