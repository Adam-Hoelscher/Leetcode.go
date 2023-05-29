package main

import "container/heap"

type PairHeap struct {
	arr [][]int
	cmp func(_, _ []int) bool
}

func NewPairHeap(arr [][]int, cmp func(_, _ []int) bool) PairHeap {
	h := PairHeap{arr, cmp}
	heap.Init(&h)
	return h
}

func (h PairHeap) Len() int {
	return len(h.arr)
}

func (h PairHeap) Less(i, j int) bool {
	return h.cmp(h.arr[i], h.arr[j])
}

func (h PairHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *PairHeap) Pop() interface{} {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[:n-1]
	return x
}

func (h *PairHeap) Push(x interface{}) {
	h.arr = append(h.arr, x.([]int))
}
