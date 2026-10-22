package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}
	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		return
	}

	maxHeap := &IntHeap{}

	for range dishesCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			return
		}
		*maxHeap = append(*maxHeap, val)
	}

	heap.Init(maxHeap)

	var targetRank int
	if _, err := fmt.Scan(&targetRank); err != nil {
		return
	}

	var result int
	for range targetRank {
		popped, ok := heap.Pop(maxHeap).(int)
		if !ok {
			return
		}
		result = popped
	}

	fmt.Println(result)
}
