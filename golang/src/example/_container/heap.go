package main

import (
	"container/heap"
	"fmt"
)

type StringHeap []string

// sort.Interface
func (h StringHeap) Len() int           { return len(h) }
func (h StringHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StringHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap.Interface
func (h *StringHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}
func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapExample() {
	var _ heap.Interface = (*StringHeap)(nil)

	h := &StringHeap{"セイバー", "アーチャー", "ランサー", "キャスター", "ライダー", "アサシン", "バーサーカー"}
	heap.Init(h)
	heap.Push(h, "ギルガメッシュ先輩")
	fmt.Printf("mininum: %s\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%s ", heap.Pop(h))
	}
	fmt.Println()
}
