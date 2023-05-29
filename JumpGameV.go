package main

type IntStack []int

func (h *IntStack) Clear() {
	*h = (*h)[:0]
}

func (h *IntStack) Len() int {
	return len(*h)
}

func (h *IntStack) Peek() int {
	old := *h
	n := len(old)
	return old[n-1]
}

func (h *IntStack) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntStack) Push(x int) {
	*h = append(*h, x)
}

func maxJumps(arr []int, d int) int {

	n := len(arr)
	min := make([]int, n)
	max := make([]int, n)

	h := &IntStack{}

	for i := 0; i < n; i++ {

		for h.Len() > 0 && arr[h.Peek()] < arr[i] {
			h.Pop()
		}

		min[i] = i - d
		if h.Len() > 0 && min[i] <= h.Peek() {
			min[i] = h.Peek() + 1
		}
		if min[i] < 0 {
			min[i] = 0
		}

		h.Push(i)
	}

	h.Clear()

	for i := n - 1; i >= 0; i-- {

		for h.Len() > 0 && arr[h.Peek()] < arr[i] {
			h.Pop()
		}

		max[i] = i + d
		if h.Len() > 0 && max[i] >= h.Peek() {
			max[i] = h.Peek() - 1
		}
		if max[i] >= n {
			max[i] = n - 1
		}

		h.Push(i)
	}

	var ans int

	cache := make([]int, n)
	var rf func(i int) int
	rf = func(i int) int {
		if cache[i] > 0 {
			return cache[i]
		}
		ans := 1
		for j := min[i]; j <= max[i]; j++ {
			if j == i {
				continue
			}
			if tail := 1 + rf(j); ans < tail {
				ans = tail
			}
		}
		cache[i] = ans
		return ans
	}

	for i := 0; i < n; i++ {
		if dfs := rf(i); ans < dfs {
			ans = dfs
		}
	}

	return ans
}
