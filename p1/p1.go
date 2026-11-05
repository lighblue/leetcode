package p1

import "slices"

// Submission: 2026.11.05 | Runtime: 0 ms (beats 100%) | Memory: 5.31 MB (beats 84.35%)

func twoSum(nums []int, target int) []int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)

	rightIndex := len(sorted) - 1
	index := 0
	for index = range sorted {
		expected := target - sorted[index]
		for ; rightIndex > index; rightIndex-- {
			if sorted[rightIndex] == expected {
				return getIndex(nums, sorted[index], expected)
			}

			if sorted[rightIndex] < expected {
				break
			}
		}
	}

	return []int{}
}

func getIndex(nums []int, num1 int, num2 int) []int {
	idx1, idx2 := -1, -1
	for index, num := range nums {
		if idx1 == -1 && num == num1 {
			idx1 = index
		} else if idx2 == -1 && num == num2 {
			idx2 = index
		}
		if idx1 != -1 && idx2 != -1 {
			break
		}
	}
	return []int{idx1, idx2}
}

//func twoSum(nums []int, target int) []int {
//	result := make([]int, 2)
//	for idx, val := range nums {
//		expected := target - val
//		if index := findValue(nums[idx+1:], expected); index != -1 {
//			result[0] = idx
//			result[1] = index + idx + 1
//			return result
//		}
//	}
//	return result
//}
//
//func findValue(nums []int, target int) int {
//	for index, value := range nums {
//		if value == target {
//			return index
//		}
//	}
//	return -1
//}
