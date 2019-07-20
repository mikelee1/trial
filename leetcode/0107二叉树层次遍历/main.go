package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left:&TreeNode{
				Val:3,
				Left:&TreeNode{
					Val:4,
					Left:&TreeNode{
						Val: 5,
					},
				},
			},
		},
		//Right: &TreeNode{
		//	Val: 4,
		//	Left: &TreeNode{
		//		Val: 5,
		//	},
		//	Right: &TreeNode{
		//		Val: 6,
		//	},
		//},
	}
	fmt.Println(levelOrderBottom(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil || *root == (TreeNode{}) {
		return res
	}
	cacheNodes := []*TreeNode{root}
	hang := []int{1}
	headhang := 1
	i := 0
	for len(cacheNodes) > i {
		headNode := cacheNodes[i]

		if headNode.Left != nil {
			cacheNodes = append(cacheNodes, headNode.Left)
			hang = append(hang, headhang+1)
		}
		if headNode.Right != nil {
			cacheNodes = append(cacheNodes, headNode.Right)
			hang = append(hang, headhang+1)
		}
		if len(res) < hang[i] {
			res = append([][]int{[]int{headNode.Val}},res...)
		} else {
			res[0] = append(res[0], headNode.Val)
		}
		if headhang == 1||(i+1 <len(hang) && headhang < hang[i+1]) {
			headhang++
		}
		i++
	}
	return res
}
