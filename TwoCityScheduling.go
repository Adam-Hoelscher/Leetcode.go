package main

import (
	"container/heap"
)

func twoCitySchedCost(costs [][]int) int {

	aHeap := NewPairHeap(nil, func(i, j []int) bool {
		return i[0]-i[1] > j[0]-j[1]
	})

	bHeap := NewPairHeap(nil, func(i, j []int) bool {
		return i[0]-i[1] < j[0]-j[1]
	})

	for _, x := range costs {
		if x[0] < x[1] {
			aHeap.Push(x)
		} else {
			bHeap.Push(x)
		}
	}

	heap.Init(&aHeap)
	for aHeap.Len() > bHeap.Len() {
		bHeap.Push(heap.Pop(&aHeap))
	}

	heap.Init(&bHeap)
	for aHeap.Len() < bHeap.Len() {
		aHeap.Push(heap.Pop(&bHeap))
	}

	var ans int

	for _, x := range aHeap.arr {
		ans += x[0]
	}
	for _, x := range bHeap.arr {
		ans += x[1]
	}

	return ans
}
