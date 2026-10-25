package maxheap

type IntHeap []int

func (h *IntHeap) Len() int {
	if h == nil {
		return 0
	}
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if h == nil || i < 0 || j < 0 || i >= len(*h) || j >= len(*h) {
		return false
	}
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if h == nil || i < 0 || j < 0 || i >= len(*h) || j >= len(*h) {
		return
	}
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}
	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	if h == nil || len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
