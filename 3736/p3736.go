package p736

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
