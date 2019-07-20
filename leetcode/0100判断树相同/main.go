package main

import "fmt"

func main() {
	a := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
		},
	}
	b := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
		},
		Right:&TreeNode{
			Val:3,
		},
	}
	fmt.Println(isSameTree(a,b))
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if (p == nil) != (q == nil){
		return false
	}
	if p.Left ==nil && p.Right == nil {
		if p.Val == q.Val{
			return true
		}
		return false
	}

	if p.Val != q.Val{
		return false
	}
	fmt.Println(p,q)
	return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
}