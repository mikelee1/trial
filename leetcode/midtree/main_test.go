package midtree_test

import (
	"testing"
	"myproj/try/leetcode/midtree"
	"github.com/magiconair/properties/assert"
)


func Test_main(t *testing.T) {
	var p *midtree.TreeNode
	var a,res []int
	//p = &midtree.TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &midtree.TreeNode{
	//		Val: 2,
	//		Left: &midtree.TreeNode{
	//			Val: 3,
	//		},
	//	},
	//}
	//a = []int{1,3,2}
	//res = midtree.InorderTraversal(p)
	//assert.Equal(t,res,a)


	p = &midtree.TreeNode{}
	nilTree := &midtree.TreeNode{}
	a = []int{}
	assert.Equal(t,p,nilTree)
	res = midtree.InorderTraversal(p)
	assert.Equal(t,res,a)

}
