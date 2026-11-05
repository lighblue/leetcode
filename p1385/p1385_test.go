package p1385

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		d    int
		want int
	}{
		{name: "example 1", arr1: []int{4, 5, 8}, arr2: []int{10, 9, 1, 8}, d: 2, want: 2},
		{name: "example 2", arr1: []int{1, 4, 2, 3}, arr2: []int{-4, -3, 6, 10, 20, 30}, d: 3, want: 2},
		{name: "example 3", arr1: []int{2, 1, 100, 3}, arr2: []int{-5, -2, 10, -3, 7}, d: 6, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue(tt.arr1, tt.arr2, tt.d); got != tt.want {
				t.Errorf("findTheDistanceValue(%v, %v, %d) = %d, want %d", tt.arr1, tt.arr2, tt.d, got, tt.want)
			}
		})
	}
}
