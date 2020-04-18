package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l11 := &ListNode{-9, &ListNode{3, nil}}
	l21 := &ListNode{5, &ListNode{7, nil}}
	a := mergeTwoLists2(l11, l21)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}

	//var l12 *ListNode
	//var l22 *ListNode
	//b := mergeTwoLists2(l12, l22)
	//for b != nil {
	//	fmt.Println(b.Val)
	//	b = b.Next
	//}
}

func listLength(l *ListNode) int {
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	totalres := res
	if l1 == nil && l2 == nil {
		return nil
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res.Val = l2.Val
			l2 = l2.Next

		} else {
			res.Val = l1.Val
			l1 = l1.Next
		}
		res.Next = &ListNode{}
		res = res.Next
	}

	if l1 == nil && l2 != nil {
		res.Val = l2.Val
		res.Next = l2.Next
	}
	if l2 == nil && l1 != nil {
		res.Val = l1.Val
		res.Next = l1.Next
	}
	return totalres
}
