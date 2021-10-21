package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) Create(list []int) {
	cur := node
	for i := 0; i < len(list); i++ {
		next := &ListNode{Val: list[i]}
		cur.Next = next
		cur = next
	}
}

func main() {
	dummyHead := &ListNode{}
	dummyHead.Create([]int{1, 1, 2, 2, 3, 3, 4})
	temp := dummyHead
	for temp.Next != nil {
		fmt.Println(temp.Next.Val)
		temp = temp.Next
	}
}
