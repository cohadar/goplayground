package container

import (
	"container/heap"
	"testing"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (ph *IntHeap) Push(x interface{}) {
	*ph = append(*ph, x.(int))
}

func (ph *IntHeap) Pop() interface{} {
	h := *ph
	last := len(h) - 1
	ret := h[last]
	*ph = h[:last]
	return ret
}

func TestHeap(t *testing.T) {
	h := IntHeap([]int{4, 5, 7, 0, 3, 6, 1, 2})
	heap.Init(&h)
	heap.Push(&h, 9)
	heap.Push(&h, 8)
	for i := 0; i < 10; i++ {
		if heap.Pop(&h).(int) != i {
			t.Errorf("%d: %d", i, heap.Pop(&h).(int))
		}
	}
}
