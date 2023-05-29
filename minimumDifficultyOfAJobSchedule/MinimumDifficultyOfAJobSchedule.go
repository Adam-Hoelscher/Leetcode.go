package main

import (
	"math"
)

func minDifficulty(jobDifficulty []int, D int) int {

	n := len(jobDifficulty)

	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = math.MaxInt32
	}
	cache[n] = 0

	for u := 1; u <= D; u++ {

		for i := 0; i+u <= n; i++ {
			min := math.MaxInt64
			max := math.MinInt64

			for j := i; j+u <= n; j++ {
				if max < jobDifficulty[j] {
					max = jobDifficulty[j]
				}
				if tail := max + cache[j+1]; min > tail {
					min = max + cache[j+1]
				}
			}

			cache[i] = min
		}

	}

	return cache[0]
}

func minDifficultyRF(jobDifficulty []int, d int) int {

	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	cache := make([][]*int, n)
	for i := range cache {
		cache[i] = make([]*int, d+1)
	}

	for i, max := n-1, 0; i >= 0; i-- {
		if max < jobDifficulty[i] {
			max = jobDifficulty[i]
		}
		m := max
		cache[i][1] = &m
	}

	var rf func(i, d int) int
	rf = func(i, d int) int {

		if cache[i][d] != nil {
			return *cache[i][d]
		}

		min := math.MaxInt
		max := math.MinInt

		for j := i; j+d <= n; j++ {

			if max < jobDifficulty[j] {
				max = jobDifficulty[j]
			}

			if tail := rf(j+1, d-1); min > max+tail {
				min = max + tail
			}
		}

		cache[i][d] = &min
		return min
	}

	return rf(0, d)
}
