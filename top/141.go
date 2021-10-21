package main

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = head.Next.Next
		slow = head.Next
		if fast == slow {
			return true
		}
	}
	return false
}
