package p1385

import "slices"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)
	arr2Min := arr2[0]
	arr2Max := arr2[len(arr2)-1]

	count := 0
	for _, num1 := range arr1 {
		minNum := num1 - d
		maxNum := num1 + d
		if minNum > arr2Max || maxNum < arr2Min {
			count++
			continue
		}

		if isValid(minNum, maxNum, arr2) {
			count++
		}
	}
	return count
}

func isValid(minNum int, maxNum int, arr []int) bool {
	for _, num2 := range arr {
		if num2 > maxNum {
			return true
		}
		if num2 >= minNum && num2 <= maxNum {
			break
		}
	}
	return false
}
