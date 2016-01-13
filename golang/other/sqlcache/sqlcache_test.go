package sqlcache

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp"
)

var testQuery = "SELECT * FROM quest WHERE id in (?)"
var testArgs = []interface{}{
	2,
}

func TestQuery(t *testing.T) {
	s := New()
	defer s.Close()
	q := s.Query(testQuery, testArgs...)
	pp.Println(q)
}

func TestQueryKVSCache(t *testing.T) {
	s := New()
	defer s.Close()
	cacheKey := fmt.Sprintf("%s:%v", testQuery, testArgs)
	s.KVS.Del(cacheKey)

	q := s.QueryKVSCache(testQuery, testArgs...)
	pp.Println(q)

	q = s.QueryKVSCache(testQuery, testArgs...)
	pp.Println(q)
}

func TestQueryMemoryCache(t *testing.T) {
	s := New()
	defer s.Close()
	q := s.QueryMemoryCache(testQuery, testArgs...)
	pp.Println(q)

	q = s.QueryMemoryCache(testQuery, testArgs...)
	pp.Println(q)
}

func BenchmarkQuery___________(b *testing.B) {
	s := New()
	defer s.Close()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Query(testQuery, testArgs...)
	}
}

func BenchmarkQueryKVSCache___(b *testing.B) {
	s := New()
	defer s.Close()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.QueryKVSCache(testQuery, testArgs...)
	}
}

func BenchmarkQueryMemoryCache(b *testing.B) {
	s := New()
	defer s.Close()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.QueryMemoryCache(testQuery, testArgs...)
	}
}
