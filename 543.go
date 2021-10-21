package main

//func main() {
//
//}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	checkDiameter(root, &res)
	return res
}

func checkDiameter(root *TreeNode, res *int) int {
	left := checkDiameter(root.Left, res)
	right := checkDiameter(root.Right, res)
	*res = max(*res, left+right)
	return max(left, right) + 1
}
