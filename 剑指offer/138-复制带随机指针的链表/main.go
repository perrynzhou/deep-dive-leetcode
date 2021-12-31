package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	fmt.Println("-----copyRandomList--------")

	nodeMapping := make(map[*Node]*Node)
	var newHead, oriNode, curNode *Node
	oriNode = head
	for {
		if oriNode == nil {
			break
		}
		tmpNode := &Node{
			Val:  oriNode.Val,
			Next: nil,
		}
		nodeMapping[oriNode] = tmpNode
		if newHead == nil {
			newHead = tmpNode
			curNode = tmpNode
		} else {
			curNode.Next = tmpNode
			curNode = tmpNode
		}
		oriNode = oriNode.Next
	}
	oriNode = head
	curNode = newHead
	for {
		if oriNode == nil {
			break
		}
		curNode.Random = nodeMapping[oriNode.Random]
		if curNode.Next == nil {
			fmt.Printf("newNode=%v,next=%v,randdom=%v\n", curNode.Val, nil, curNode.Random.Val)
		} else {
			fmt.Printf("newNode=%v,next=%v,randdom=%v\n", curNode.Val, curNode.Next.Val, curNode.Random.Val)
		}
		oriNode = oriNode.Next
		curNode = curNode.Next
	}
	return newHead
}

func buildNode(n int) *Node {
	var newHead, curNode *Node
	var randIndexMapping []int
	nodeIndexMapping := make(map[int]*Node)
	i := 0
	for {
		if i >= n {
			break
		}
		tmpNode := &Node{
			Val:  rand.Intn(100),
			Next: nil,
		}
		if newHead == nil {
			newHead = tmpNode
			curNode = tmpNode
		} else {
			curNode.Next = tmpNode
			curNode = tmpNode
		}
		randIndexMapping = append(randIndexMapping, i)
		nodeIndexMapping[i] = curNode
		i++
	}
	// shuffer index
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randIndexMapping), func(i, j int) {
		randIndexMapping[i], randIndexMapping[j] = randIndexMapping[j], randIndexMapping[i]
	})
	fmt.Println("-----shffle random index--------")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", randIndexMapping[i])
	}

	fmt.Println()
	curNode = newHead
	i = 0
	for {
		if i >= n {
			break
		}
		index := randIndexMapping[i]
		curNode.Random = nodeIndexMapping[index]
		i++
		if curNode.Next == nil {
			fmt.Printf("newNode=%v,next=%v,randdom=%v\n", curNode.Val, nil, curNode.Random.Val)
		} else {
			fmt.Printf("newNode=%v,next=%v,randdom=%v\n", curNode.Val, curNode.Next.Val, curNode.Random.Val)
		}

		curNode = curNode.Next

	}
	return newHead
}
func main() {
	oriNode := buildNode(4)
	copyRandomList(oriNode)
}
