package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


func main()  {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:3,
			},
		},
	}
	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
			},
		},
	}
	l3 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
			},
		},
	}
	a := mergeKLists(l1,l2,l3)
	for a != nil{
		time.Sleep(1*time.Second)
		fmt.Println("a:",a.Val)
		a = a.Next
	}

}
func mergeKLists(ls ...*ListNode) *ListNode {
	fmt.Println(ls)
	var head,node *ListNode
	var secondmin int
	//var firstmin int
	//var totalnode int
	newls := []*ListNode{}
	for _,l := range ls{
		if l != nil{
			newls = append(newls,l)
		}
	}
	if len(newls)==0{
		return nil
	}
	head = newls[0]
	node = newls[0]
	for _,v := range newls{
		if v.Val < node.Val{
			node = v
			//firstmin = v.Val
		}
		if node.Val < secondmin{
			secondmin = node.Val
		}
	}



	return head
}
