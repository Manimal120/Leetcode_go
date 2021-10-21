package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node0 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node0.Left = nil
	node0.Right = &node2
	node2.Left = &node3
	fmt.Println(threeOrders(&node0))
}

func threeOrders(root *TreeNode) [][]int {

	var res = [][]int{{}, {}, {}}
	preOrder(root, &res[0])
	inOrder(root, &res[1])
	postOrder(root, &res[2])
	return res
}

func preOrder(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preOrder(root.Left, result)
		preOrder(root.Right, result)
	}
}

func inOrder(root *TreeNode, result *[]int) {
	if root != nil {
		inOrder(root.Left, result)
		*result = append(*result, root.Val)
		inOrder(root.Right, result)
	}
}

func postOrder(root *TreeNode, result *[]int) {
	if root != nil {
		postOrder(root.Left, result)
		postOrder(root.Right, result)
		*result = append(*result, root.Val)
	}
}
