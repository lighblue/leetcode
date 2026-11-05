package p736

// Submission: 2026.11.03 | Runtime: 0 ms (beats 100%) | Memory: 4.76 MB (beats 95.42%)

func minMoves(nums []int) int {
	max, count := 0, 0
	for index, num := range nums {
		if num <= max {
			count += max - num
			continue
		}
		count += index * (num - max)
		max = num
	}

	return count
}
