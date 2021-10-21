package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	pre := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			pre = pre.Next
		}
		i++
	}
	pre.Next = pre.Next.Next
	return dummyHead.Next
}
