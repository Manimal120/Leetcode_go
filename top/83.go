package main

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	if head == nil || head.Next == nil {
		return head
	}

	for curr.Next != nil {
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
