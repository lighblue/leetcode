package p462

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
