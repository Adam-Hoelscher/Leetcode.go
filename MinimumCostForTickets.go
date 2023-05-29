package main

import "sort"

func mincostTickets(days []int, costs []int) int {

	n := len(days)

	cache := make([]int, n)
	var rf func(_ int) int
	rf = func(i int) int {

		if i == n {
			return 0
		}

		if cache[i] > 0 {
			return cache[i]
		}

		// start off assuming we buy 1 day
		ans := costs[0] + rf(i+1)

		// check if buying 7 is better
		if x := costs[1] + rf(sort.SearchInts(days, days[i]+7)); ans > x {
			ans = x
		}

		// check if buying 30 is better
		if x := costs[2] + rf(sort.SearchInts(days, days[i]+30)); ans > x {
			ans = x
		}

		cache[i] = ans
		return ans
	}

	return rf(0)
}
