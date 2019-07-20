package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	var carry = 0

	for l1 != nil || l2 != nil {
		fmt.Println(l1.Val, l2.Val)

		tmpnode := &ListNode{}
		y, x := addTwo(l1.Val, l2.Val, carry)
		carry = y
		l3.Val = x
		l3.Next = tmpnode
		l3 = l3.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}
	return l3
}

func addTwo(a, b, carry int) (int, int) {
	c := a + b + carry
	x := c % 10
	y := (c - x) / 10
	return y, x
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("%#v", l3)
}
