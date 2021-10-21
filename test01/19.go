package main

// solution:
// cur go first, the distance is 1
// when cur go far beyond n, i>n, prev move to the next
// keep moving until cur reaches the bound
// now, the prev is the target
// pre.next = pre.next.next

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
