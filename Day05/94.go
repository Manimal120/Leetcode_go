package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderOperate(root, &res)
	return res
}

func inorderOperate(root *TreeNode, res *[]int) {
	if root != nil {
		inorderOperate(root.Left, res)
		*res = append(*res, root.Val)
		inorderOperate(root.Right, res)
	}
}
