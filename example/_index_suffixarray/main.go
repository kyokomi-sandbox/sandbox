package main

import (
	"fmt"
	"index/suffixarray"
	"sort"

	"github.com/cloudflare/ahocorasick"
	trigram "github.com/dgryski/go-trigram"
)

func main() {
	fmt.Println(_ahocorasick())
	fmt.Println(_indexSuffixArray())
	fmt.Println(_trigram())
}

func _ahocorasick() string {
	patterns := []string{
		"mercury", "venus", "earth", "mars",
		"jupiter", "saturn", "uranus", "pluto",
	}

	m := ahocorasick.NewStringMatcher(patterns)

	found := m.Match([]byte(`earth`))
	return fmt.Sprintln("found patterns", found)
}

func _indexSuffixArray() string {
	docs := []string{
		"mercury", "venus", "earth", "mars",
		"jupiter", "saturn", "uranus", "pluto",
	}

	var data []byte
	var offsets []int

	for _, d := range docs {
		data = append(data, []byte(d)...)
		offsets = append(offsets, len(data))
	}
	sfx := suffixarray.New(data)

	query := "earth"

	idxs := sfx.Lookup([]byte(query), -1)
	var results []int
	for _, idx := range idxs {
		i := sort.Search(len(offsets), func(i int) bool { return offsets[i] > idx })
		if idx+len(query) <= offsets[i] {
			results = append(results, i)
		}
	}

	return fmt.Sprintf("%q is in documents %v\n", query, results)
}

func _trigram() string {
	docs := []string{
		"mercury", "venus", "earth", "mars",
		"jupiter", "saturn", "uranus", "pluto",
	}

	idx := trigram.NewIndex(docs)

	found := idx.Query("earth")
	return fmt.Sprintln("matched documents", found)
}
