package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	list1Count = flag.Int("l1", 2, "list1 length")
	list2Count = flag.Int("l2", 3, "list2 length")
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// l2={2,3,4},l2={5,6,4}
func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("#############addTwoNumber#############")
	var nodeHead *ListNode
	var tmpNode *ListNode 
	increment := 0
	for {
		tmpSum := 0
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			tmpSum = tmpSum + l1.Val
			l1 = l1.Next

		}
		if l2 != nil {
			tmpSum = tmpSum + l2.Val
			l2 = l2.Next
		}
		tmpSum = increment + tmpSum
		var val int
		if tmpSum > 9 {
			increment = tmpSum / 10
			val = tmpSum % 10
		} else {
			increment = 0
			val = tmpSum
		}
		tmp := &ListNode{
			Val:  val,
			Next: nil,
		}
		if nodeHead == nil {
			nodeHead = tmp
			tmpNode = tmp
		} else {
			tmpNode.Next = tmp
			tmpNode = tmp
		}

	}
	if increment > 0 {
		tmp := &ListNode{
			Val:  increment,
			Next: nil,
		}
		tmpNode.Next = tmp
		tmpNode = tmp
	}
	node := nodeHead
	for {
		fmt.Printf("%v ", node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	return nodeHead
}

func initListNode(n int,val int) *ListNode {

	var tmpNode *ListNode
	for i := 0; i < n; i++ {
		tmpVal  := 0
		if tmpVal <0 {
			tmpVal = rand.Intn(10)
		} else {
			tmpVal = val
		}
		tmp := &ListNode{
			Val:  tmpVal,
			Next: tmpNode,
		}
		tmpNode = tmp
	}
	fmt.Println("##########################")
	node := tmpNode
	for {
		fmt.Printf("%v ", node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	fmt.Println()
	return tmpNode
}
func main() {
	flag.Parse()
	l1 := initListNode(*list1Count,9)
	l2 := initListNode(*list2Count,9)
	l3 := addTwoNumber(l1, l2)
	fmt.Println(l3)
}
