package p462

// Submission: 2026.11.02 | Runtime: 0 ms (beats 100%) | Memory: 6.19 MB (beats 85.14%)

import (
	"slices"
)

func minMoves2(nums []int) int {
	slices.Sort(nums)
	midNum := nums[len(nums)/2]
	count := 0
	for _, num := range nums[:len(nums)/2] {
		count += midNum - num
	}

	for _, num := range nums[len(nums)/2:] {
		count += num - midNum
	}

	return count
}
