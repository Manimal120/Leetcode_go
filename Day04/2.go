package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1, n2 int
	var carry int
	dummyHead := &ListNode{}
	current := dummyHead
	for l1 != nil || l2 != nil || carry != 0 {
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
			Val:  (n1 + n2 + carry) % 10,
			Next: nil,
		}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return dummyHead.Next
}
