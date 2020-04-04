package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var data = TreeNode{1,
	&TreeNode{2,
		&TreeNode{4, &TreeNode{7, nil, nil}, nil},
		&TreeNode{5, nil, nil}},
	&TreeNode{3, nil,
		&TreeNode{6, nil, &TreeNode{8, nil, nil}}},
}

func main() {
	fmt.Println(deepestLeavesSum(&data))
}

type Node struct {
	*TreeNode
	Row    int
	Column int
}

func deepestLeavesSum(root *TreeNode) int {
	var cache = map[string]int{}
	var nodes []*Node
	nodes = append(nodes, &Node{root, 1, 1})
	var res int
	var head *Node
	for len(nodes) > 0 {
		head = nodes[0]
		if head.Left != nil {
			nodes = append(nodes, &Node{head.Left, head.Row + 1, head.Column*2 - 1})
		}
		if head.Right != nil {
			nodes = append(nodes, &Node{head.Right, head.Row + 1, head.Column * 2})
		}
		key := strconv.Itoa(head.Row) + "-" + strconv.Itoa(head.Column)
		nodes = nodes[1:]
		cache[key] = head.Val
	}
	for k, v := range cache {
		if strings.Contains(k, strconv.Itoa(head.Row)+"-") {
			res += v
		}
	}

	return res
}
