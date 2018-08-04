package main

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.New()
	tokens := t.Tokenize("MMOで複数人のユーザーが会話したり、コミュニケーションをとれる空間。")
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		if token.Features()[0] != "名詞" {
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
