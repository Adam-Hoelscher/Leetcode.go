package main

import "sort"

func nextPermutation(nums []int) {

	n := len(nums)

	// Find the longest tail of non-increasing values at the end of the array
	tailBeg := n - 1
	for tailBeg > 0 && nums[tailBeg-1] >= nums[tailBeg] {
		tailBeg--
	}

	// swap the tail front-to-back; it is now sorted
	for i, j := tailBeg, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	if tailBeg == 0 {
        return
	}

    // if the tail was not the entire array, we need to swap the pivot into the
	// tail. the new value at the pivot position needs to be the lowest value
	// that is greater than the current pivot
    pivot := tailBeg + sort.SearchInts(nums[tailBeg:], nums[tailBeg-1]+1)
    nums[tailBeg-1], nums[pivot] = nums[pivot], nums[tailBeg-1]

}
