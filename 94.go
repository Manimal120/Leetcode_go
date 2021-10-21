package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//	node0 := TreeNode{1, nil, nil}
//	node2 := TreeNode{2, nil, nil}
//	node3 := TreeNode{3, nil, nil}
//	node0.Left = nil
//	node0.Right = &node2
//	node2.Left = &node3
//	fmt.Println(inorderTraversal(&node0))
//}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
