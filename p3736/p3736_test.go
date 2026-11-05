package p736

import "testing"

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "decrease then increase", nums: []int{2, 1, 3}, want: 3},
		{name: "equal then increase", nums: []int{4, 4, 5}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.nums); got != tt.want {
				t.Errorf("minMoves(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
