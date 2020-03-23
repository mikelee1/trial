package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var h = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
	},
}

func main() {
	fmt.Println(middleNode(h))
}

func middleNode(head *ListNode) *ListNode {
	length := 1
	resNode := head
	for head.Next != nil {
		length++
		head = head.Next
		if length%2 == 0 {
			resNode = resNode.Next
		}
	}
	return resNode
}
