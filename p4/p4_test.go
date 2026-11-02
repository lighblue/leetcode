package p4

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{name: "odd total length", nums1: []int{1, 3}, nums2: []int{2}, want: 2},
		{name: "even total length", nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.5},
		{name: "first array empty", nums1: []int{}, nums2: []int{1}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
