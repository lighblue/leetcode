package p462

import "testing"

func TestMinMoves2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "three numbers", nums: []int{1, 2, 3}, want: 2},
		{name: "four numbers", nums: []int{1, 2, 9, 10}, want: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.nums); got != tt.want {
				t.Errorf("minMoves2(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
