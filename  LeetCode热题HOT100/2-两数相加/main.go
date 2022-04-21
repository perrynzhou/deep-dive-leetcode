package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var incr int
	var head, tail *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}
		var value int
		if l1 != nil {
			value = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}

		value += incr
		if value >= 10 {
			incr = 1
		} else {
			incr = 0
		}
		value = value % 10
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

	}
	if incr > 0 {
		node := &ListNode{
			Val:  1,
			Next: nil,
		}
		tail.Next = node
		tail = node
	}
	return head
}

func NewList(arr []int) *ListNode {
	var head, tail *ListNode
	for _, v := range arr {
		tmp := &ListNode{
			Val:  v,
			Next: nil,
		}
		if head == nil {
			head = tmp
			tail = tmp
		} else {
			tail.Next = tmp
			tail = tmp
		}
	}
	return head
}
func OutList(head *ListNode) {
	for head != nil {
		fmt.Printf("node=%p,value=%d\n", head, head.Val)
		head = head.Next
	}
}
func main() {
	// l1 := NewList([]int{2, 3, 4})

	l1 := NewList([]int{0})
	OutList(l1)
	fmt.Println("********")
	// l2 := NewList([]int{9, 5, 6})

	l2 := NewList([]int{0})
	OutList(l2)
	fmt.Println("********")

	fmt.Println("##########")
	l3 := addTwoNumbers(l1, l2)
	OutList(l3)

}
