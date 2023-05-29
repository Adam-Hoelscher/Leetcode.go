package main

import (
	"container/heap"
	"math/bits"
)

type SearchHeap [][]int

func (h SearchHeap) Len() int {
	return len(h)
}

func (h SearchHeap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}

func (h SearchHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SearchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *SearchHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {

	m := len(students)
	n := len(students[0])
	mask := uint(1)<<n - 1

	makeScores := func(people [][]int) []uint {
		score := make([]uint, m)
		for i, p := range people {
			for j, resp := range p {
				score[i] |= uint(resp) << j
			}
		}
		return score
	}

	sScores := makeScores(students)
	mScores := makeScores(mentors)

	queue := new(SearchHeap)
	queue.Push([]int{0, 0, 0})

	for queue.Len() > 0 {

		// get the best state from the queue and unpack it
		state := heap.Pop(queue).([]int)
		loss, i, taken := state[0], state[1], state[2]

		// if we've assigned all the students, terminate
		if i == m {
			return m*n - loss
		}

		// for each mentor
		for j := 0; j < m; j++ {
			// if that mentors is taken, skip them
			if (taken>>j)&1 == 1 {
				continue
			}

			// calculate how much farther we get from perfect for this mentor
			mismatch := bits.OnesCount(mask & (sScores[i] ^ mScores[j]))
			nextLoss := loss + mismatch

			// tag the mentor as taken
			nextTaken := taken | (1 << j)

			// add the state where this student gets this mentor to the queue
			nextState := []int{nextLoss, i + 1, nextTaken}
			heap.Push(queue, nextState)
		}
	}

	return -1
}
