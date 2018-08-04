package main

import "testing"

/*
$ go test --bench .
testing: warning: no tests to run
PASS
BenchmarkSqlxNamed___________________	  200000	      6053 ns/op	  46.09 MB/s	    1136 B/op	      13 allocs/op
BenchmarkSqlxNamedReplaceQueryComment	   30000	     50177 ns/op	   5.56 MB/s	    8592 B/op	     108 allocs/op
ok  	github.com/kyokomi-sandbox/go-sandbox/libexample/_sqlx	3.290s
*/

func BenchmarkSqlxNamed___________________(b *testing.B) {
	b.SetBytes(int64(len(queryString)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sqlxNamed(queryString, argsMap)
	}
}

func BenchmarkSqlxNamedReplaceQueryComment(b *testing.B) {
	b.SetBytes(int64(len(queryString)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sqlxNamedRegexpReplaceQueryComment(queryString, argsMap)
	}
}
