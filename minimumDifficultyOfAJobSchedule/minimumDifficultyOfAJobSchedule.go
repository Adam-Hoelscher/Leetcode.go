package main

import (
	"math"
)

func minDifficulty(jobDifficulty []int, d int) int {

	n := len(jobDifficulty)

	if n < d {
		return -1
	}

	cache := make([]int, n+1)
	for i := range cache {
		cache[i] = math.MaxInt32
	}
	cache[n] = 0

	for u := 1; u <= d; u++ {

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
