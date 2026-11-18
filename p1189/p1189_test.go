package p1189

import "testing"

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{name: "example 1", text: "nlaebolko", want: 1},
		{name: "example 2", text: "loonbalxballpoon", want: 2},
		{name: "example 3", text: "leetcode", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons(%q) = %d, want %d", tt.text, got, tt.want)
			}
		})
	}
}
