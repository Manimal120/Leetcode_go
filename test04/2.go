package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	n1, n2, carry, current := 0, 0, 0, dummyHead

	for l1 != nil || l2 != nil {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{
			Val: (n1 + n2 + carry) % 10,
		}
		carry = (n1 + n2) / 10
		current = current.Next
	}
	return dummyHead.Next
}
