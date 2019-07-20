package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:5,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
			},
		},
	}
	l2 := &ListNode{
		Val:5,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
			},
		},
	}
	a := addTwoNumbers(l1,l2)
	for a!=nil{
		fmt.Println(a.Val)
		a = a.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	node := res
	sum := 0
	for l1!=nil||l2!=nil||sum != 0{
		if l1==nil&&l2==nil{
			node.Val = sum
			break
		}
		if l1 == nil{
			node.Val,sum = addonce(l2.Val,0,sum)
			l2 = l2.Next
			if sum==0&&l2==nil{
				break
			}
			node.Next = &ListNode{}
			node = node.Next

			continue
		}
		if l2 == nil{
			node.Val,sum = addonce(l1.Val,0,sum)
			l1 = l1.Next
			if sum==0&&l1==nil{
				break
			}
			node.Next = &ListNode{}
			node = node.Next
			continue
		}

		node.Val,sum = addonce(l1.Val,l2.Val,sum)
		l1 = l1.Next
		l2 = l2.Next
		if sum==0&&l1==nil&&l2 == nil{
			break
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	return res
}

func addonce(a,b,sum int) (int,int) {
	return (sum + a+b) % 10,(sum + a+b) / 10
}