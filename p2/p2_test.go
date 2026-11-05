package p2

import (
	"slices"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{name: "example 1", l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
		{name: "example 2", l1: []int{0}, l2: []int{0}, want: []int{0}},
		{name: "example 3", l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}, want: []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{name: "unequal list lengths", l1: []int{1, 8}, l2: []int{0}, want: []int{1, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listValues(addTwoNumbers(makeList(tt.l1), makeList(tt.l2)))
			if !slices.Equal(got, tt.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}

func makeList(values []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return dummy.Next
}

func listValues(head *ListNode) []int {
	var values []int
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}
