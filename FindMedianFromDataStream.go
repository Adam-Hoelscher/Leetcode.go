package main

import (
	"container/heap"
)

type Heap struct {
	array  []interface{}
	before func(i, j interface{}) bool
}

func NewHeap(Before func(i, j interface{}) bool) *Heap {
	h := new(Heap)
	h.before = Before
	return h
}

func (h Heap) Len() int {
	return len(h.array)
}

func (h Heap) Less(i, j int) bool {
	return h.before(h.array[i], h.array[j])
}

func (h Heap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h Heap) Peek() interface{} {
	return h.array[0]
}

func (h *Heap) Pop() interface{} {
	old := h.array
	n := len(old)
	val := old[n-1]
	h.array = old[:n-1]
	return val
}

func (h *Heap) Push(x interface{}) {
	h.array = append(h.array, x)
}

// invariant: below.Len() <= above.Len()
// invariant: max(below) <= median <= min(above)
type MedianFinder struct {
	above *Heap
	below *Heap
}

func Constructor() MedianFinder {

	above := NewHeap(func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})

	below := NewHeap(func(i, j interface{}) bool {
		return i.(int) > j.(int)
	})

	return MedianFinder{
		above: above,
		below: below,
	}
}

func (this *MedianFinder) AddNum(num int) {

	heap.Push(this.below, num)

	// move from below to above as long as they are unbalanced
	if this.below.Len() > this.above.Len() {
		heap.Push(this.above, heap.Pop(this.below))
	}

	// move from above to below as long as they are unbalanced
	if this.below.Len() < this.above.Len() {
		heap.Push(this.below, heap.Pop(this.above))
	}

}

func (this *MedianFinder) FindMedian() float64 {

	// the heaps are now close to balanced. below might be one larger than above
	ans := float64(this.below.Peek().(int))

	// if below and above are the same size, then we need to interpolate
	if this.below.Len() == this.above.Len() {
		ans += float64(this.above.Peek().(int))
		ans /= 2
	}

	return ans
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
