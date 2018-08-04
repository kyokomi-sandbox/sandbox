package main

import "testing"

func BenchmarkAhocorasick_____(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ahocorasick()
	}
}

func BenchmarkIndexSuffixArray(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_indexSuffixArray()
	}
}

func BenchmarkTrigram_________(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_trigram()
	}
}
