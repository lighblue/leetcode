package p2

type ListNode struct {
	Val  int
	Next *ListNode
}

// Submission: 2026.11.05 | Runtime: 0 ms (beats 100%) | Memory: 5.8 MB (beats 98.05%)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rst := new(ListNode)
	rstH := rst

	sum, bubble := 0, 0
	reuse := new(ListNode)
	for l1 != nil || l2 != nil || bubble != 0 {
		sum = bubble
		var nextReuse *ListNode
		if l1 != nil {
			sum += l1.Val
			nextReuse = l1
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			nextReuse = l2
			l2 = l2.Next
		}
		reuse.Val = sum % 10
		reuse.Next = nil
		rst.Next = reuse
		rst = rst.Next
		bubble = sum / 10
		reuse = nextReuse
		if reuse == nil {
			reuse = new(ListNode)
		}
	}

	return rstH.Next
}
