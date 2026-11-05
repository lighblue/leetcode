package p2

type ListNode struct {
	Val  int
	Next *ListNode
}

// Submission: 2026.11.05 | Runtime: 0 ms (beats 100%) | Memory: 5.99 MB (beats 89.04%)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rst := new(ListNode)
	rstH := rst

	sum, bubble := 0, 0
	for l1 != nil || l2 != nil || bubble != 0 {
		sum = bubble
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		rst.Next = &ListNode{sum % 10, nil}
		rst = rst.Next
		bubble = sum / 10
	}

	return rstH.Next
}
