package main

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkStringReplacer(b *testing.B) {
	b.ResetTimer()

	comment := "あいうえおかきくけこさしすせそ殺すたちつてとなにぬねの"
	for i := 0; i < b.N; i++ {
		replacer.Replace(comment)
	}
}

func BenchmarkStringReplacer_Reverse(b *testing.B) {
	b.ResetTimer()

	comment := "あいうえおかきくけこさしすせそ殺すたちつてとなにぬねの"
	for i := 0; i < b.N; i++ {
		replaceComment := replacer.Replace(comment)
		regexp.MustCompile(strings.Replace(replaceComment, "***", `([\s\S]*)`, -1)).FindStringSubmatch(comment)
	}
}

func BenchmarkStringReplacerHit(b *testing.B) {
	b.ResetTimer()

	comment := "あいうえおかきくけこさしすせそ殺すたちつてとなにぬねの"
	for i := 0; i < b.N; i++ {
		StringReplacerHit(comment)
	}
}

func BenchmarkSearchStringIndex(b *testing.B) {
	b.ResetTimer()

	comment := "あいうえおかきくけこさしすせそ殺すたちつてとなにぬねの"
	for i := 0; i < b.N; i++ {
		SearchStringIndex(comment)
	}
}

func BenchmarkStringRegexpMatchString(b *testing.B) {
	b.ResetTimer()

	comment := "あいうえおかきくけこさしすせそ殺すたちつてとなにぬねの"
	for i := 0; i < b.N; i++ {
		reg.MatchString(comment)
	}
}
