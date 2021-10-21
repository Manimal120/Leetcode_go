package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 1. 干嘛
// 2. 结束
// 3. 如何进行下一个 || 如何缩小范围
//缩小之后，我们可以通过一些辅助的变量或者操作，使原函数的结果不变。
//例如，f(n) 这个范围比较大，我们可以让 f(n) = n * f(n-1)。这样，
//范围就由 n 变成了 n-1 了，范围变小了，并且为了原函数f(n) 不变，我们需要让 f(n-1) 乘以 n。
