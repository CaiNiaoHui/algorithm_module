package _package

// 先序非递归 通过stack实现
func Preorder_non_recursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
 		return result
	}
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Value)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}