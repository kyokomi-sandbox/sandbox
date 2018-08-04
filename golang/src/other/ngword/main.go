package main

import (
	"fmt"
	"regexp"
	"strings"
)

var NGWords2 = []string{
	"寿司",
	"ピザ",
}

// https://docs.google.com/spreadsheets/d/1mOfBuW8vW1wT1NIYH_Enong0OKD5WEwoiRAs2ZH5I7o/edit#gid=0

var reg = regexp.MustCompile(strings.Join(NGWords2, "|"))
var replacer *strings.Replacer
var NGWords []string

func init() {
	for i := range NGWords2 {
		NGWords = append(NGWords[:(i*2)], append([]string{NGWords2[i], "***"}, NGWords[(i*2):]...)...)
	}
	replacer = strings.NewReplacer(NGWords...)
}

func SearchStringIndex(text string) bool {
	hit := false
	for j := range NGWords2 {
		if index := strings.Index(text, NGWords2[j]); index != -1 {
			hit = true
			break
		}
	}
	return hit
}

func StringReplacerHit(text string) bool {
	return replacer.Replace(text) != text
}

func main() {
	ngComment := "あいうピザえおかきくけこさしす寿司たちつてとなにぬピザねの"
	fmt.Println("Before:", ngComment)
	fmt.Println("After:", replacer.Replace(ngComment))
	fmt.Println("After:", StringReplacerHit(ngComment))
	fmt.Println("After:", SearchStringIndex(ngComment))
	fmt.Println("After:", reg.MatchString(ngComment))

	format := strings.Replace(replacer.Replace(ngComment), "***", `([\s\S]*)`, -1)
	res := regexp.MustCompile(format).FindStringSubmatch(ngComment)
	fmt.Println(res)

	okComment := "あいうえおかきくけこさしすせそたちつてとなにぬねの"
	fmt.Println("Before:", okComment)
	fmt.Println("After:", replacer.Replace(okComment))
	fmt.Println("After:", StringReplacerHit(okComment))
	fmt.Println("After:", SearchStringIndex(okComment))
	fmt.Println("After:", reg.MatchString(okComment))
}
