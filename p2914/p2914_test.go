package p2914

import "testing"

func TestMinChanges(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "1001", want: 2},
		{name: "example 2", s: "10", want: 1},
		{name: "example 3", s: "0000", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.s); got != tt.want {
				t.Errorf("minChanges(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
