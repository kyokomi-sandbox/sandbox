package main

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"fmt"
	"html"
)

var input = `
# Title

## Hogehoge

fugafuga

[google](http://google.com)

<a href="hoge"></a>

![img](https://camo.githubusercontent.com/f52e5ecfc8a1c3475a1f2c4173ef32940528ecb7/68747470733a2f2f7472617669732d63692e6f72672f72757373726f73732f626c61636b6672696461792e7376673f6272616e63683d6d6173746572)
`

func main() {
	esInput := html.EscapeString(input)
	unsafe := blackfriday.MarkdownCommon([]byte(esInput))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	fmt.Println(string(html))
}
