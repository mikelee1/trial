package midtree

import "github.com/op/go-logging"

var Logger *logging.Logger

func init() {
	Logger = logging.MustGetLogger("test")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
var nilNode = &TreeNode{}

func InorderTraversal(root *TreeNode) []int {
	if root == (&TreeNode{}) || root == nil {
		Logger.Info("equal")
		return []int{}
	}
	Logger.Info("one")
	if root.Left != nil {
		InorderTraversal(root.Left)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		InorderTraversal(root.Right)
	}
	defer func() {

	}()
	return res
}
