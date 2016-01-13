package main_test

import (
	"log"
	"testing"
	"time"

	"github.com/kyokomi/expcache"
	"github.com/patrickmn/go-cache"
)

type Hoge struct {
	Name string
	Num  int
}

var sampleData = Hoge{
	Name: "fuga",
	Num:  99999,
}

func BenchmarkGoCache(b *testing.B) {
	c := cache.New(5*time.Minute, 30*time.Second)

	c.Set("foo", sampleData, cache.DefaultExpiration)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if ch, ok := c.Get("foo"); ok {
			x := ch.(Hoge)
			if x.Name != sampleData.Name || x.Num != sampleData.Num {
				log.Println("miss")
			}
		}
	}
}

type exampleData struct {
	Expire expcache.ExpireOnMemoryCache
	cache  Hoge
}

type ExampleData interface {
	Get(now time.Time) (Hoge, error)
}

func NewExampleDataCache() ExampleData {
	e := &exampleData{cache: Hoge{}}
	e.Expire = expcache.NewExpireMemoryCache(e, 15*time.Minute)
	return e
}

func (e *exampleData) Get(now time.Time) (Hoge, error) {
	var dataCopy Hoge
	if err := e.Expire.ExpireRefresh(now); err != nil {
		return dataCopy, err
	}

	e.Expire.RLocker().Lock()
	dataCopy = e.cache
	e.Expire.RLocker().Unlock()

	return dataCopy, nil
}

func (d *exampleData) Refresh() error {
	d.cache = sampleData // TODO: DB or Redis or S3 etc...
	return nil
}

var _ expcache.OnMemoryCache = (*exampleData)(nil)
var _ ExampleData = (*exampleData)(nil)

func BenchmarkExpCache(b *testing.B) {
	e := NewExampleDataCache()

	b.ReportAllocs()
	b.ResetTimer()

	var now = time.Now()
	for i := 0; i < b.N; i++ {
		if ch, err := e.Get(now); err == nil {
			if ch.Name != sampleData.Name || ch.Num != sampleData.Num {
				log.Println("miss")
			}
		}
	}
}
