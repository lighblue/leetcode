package p4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid1, mid2 := 0, 0
	len1, len2 := len(nums1), len(nums2)
	mid1 = (len1 + len2) / 2
	mid2 = mid1
	if (len1+len2)%2 == 0 {
		mid1--
		mid2 = mid1 + 1
	}

	index := 0
	index1, index2 := 0, 0
	midNum1, midNum2 := 0, 0
	for index <= mid2 {
		midNum1 = midNum2
		if index2 >= len2 {
			midNum2 = nums1[index1]
			index1++
			index++
			continue
		}

		if index1 >= len1 {
			midNum2 = nums2[index2]
			index2++
			index++
			continue
		}

		if nums1[index1] <= nums2[index2] {
			midNum2 = nums1[index1]
			index1++
			index++
			continue
		}

		if nums2[index2] < nums1[index1] {
			midNum2 = nums2[index2]
			index2++
			index++
			continue
		}
	}

	if mid1 == mid2 {
		return float64(midNum2)
	}

	return float64(midNum2+midNum1) / 2
}
