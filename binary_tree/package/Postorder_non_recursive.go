package _package



func Postorder_non_recursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	var lastval *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// pop
		if node.Right == nil || node.Right == lastval{
			result = append(result, node.Value)
			stack = stack[:len(stack)-1]
			lastval = node
		} else {
			root = node.Right
		}
	}
	return result
}
