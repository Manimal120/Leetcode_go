package main

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	judge := make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		judge[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if judge[tmp] {
			return tmp
		}
	}
	return nil
}

// leetcode a+c+b = b+c+a

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
