package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead, curNode *ListNode
	var incr int
	var sum int
	for {
		// 9   9,9,9
		if l1 == nil && l2 == nil {
			break
		}
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += incr
		if sum > 9 {
			incr = 1
		} else {
			incr = 0
		}
		tmpNode := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		if newHead == nil {
			newHead = tmpNode
			curNode = tmpNode
		} else {
			curNode.Next = tmpNode
			curNode = tmpNode
		}
		fmt.Println("v=", sum%10)
	}
	if incr > 0 {
		tmpNode := &ListNode{
			Val:  incr,
			Next: nil,
		}
		curNode.Next = tmpNode
		curNode = tmpNode
		fmt.Println("v=", incr)
	}
	return newHead
}
