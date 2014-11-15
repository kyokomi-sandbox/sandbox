package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Is123(r rune) bool {
	if uint32(r) <= unicode.MaxLatin1 {
		switch r {
		case '1', '2', '3':
			return true
		}
	}
	return false
}

func Sample() {

	// Index

	baseText := "ほげアッシ盟約に従いアッシェンテ!"

	// 文字数じゃなくてバイト数
	fmt.Println("Count", len(baseText))
	// 前方一致の完全一致
	fmt.Println("Index", strings.Index(baseText, "アッシェンテ"))
	// 後方一致の完全一致
	fmt.Println("LastIndex", strings.LastIndex(baseText, "アッシェンテ"))
	// 前方一致の部分一致（この場合、「ア」の位置）
	fmt.Println("IndexAny", strings.IndexAny(baseText, "アッシェンテ"))
	// 前方一致の部分一致（この場合、「テ」の位置）
	fmt.Println("LastIndexAny", strings.LastIndexAny(baseText, "アッシェンテ"))

	// Split

	baseSplitText := "aaaa#a#bbbb#b#cccc"
	// 区切り文字が消えて配列に分かれる
	split1 := strings.Split(baseSplitText, "#")
	fmt.Println("Split", split1, len(split1))
	// 区切り文字は残してSplitする
	split2 := strings.SplitAfter(baseSplitText, "#")
	fmt.Println("SplitAfter", split2, len(split2))
	// N個区切ったら終わり
	split3 := strings.SplitN(baseSplitText, "#", 3)
	fmt.Println("SplitN", split3, len(split3))
	// N個区切ったら終わりの区切り文字残す版
	split4 := strings.SplitAfterN(baseSplitText, "#", 3)
	fmt.Println("SplitAfterN", split4, len(split4))

	// Fields
	t1 := "aaaa112bbbb3333ccccc2ddddd3"
	fmt.Println("FieldsFunc", strings.FieldsFunc(t1, Is123))

	// Join
	inputUrl1 := TrimPrefixAndSuffix("http://localhost:8080/", "/")
	inputUrl2 := TrimPrefixAndSuffix("/test/", "/")
	fmt.Println(strings.Join([]string{inputUrl1, inputUrl2}, "/"))

	// Trim
}

func TrimPrefixAndSuffix(s, fix string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, fix), fix)
}
