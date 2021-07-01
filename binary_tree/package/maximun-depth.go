package _package

// 通过分治法找出树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// divide
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// 判断平衡二叉树， 不是返回-1，是返回树的高度。  (判断树的高度的变种)

func IsBalaced(root *TreeNode) int {
	result := balaceTreeHight(root)
	return result
}

func balaceTreeHight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//divide
	left := balaceTreeHight(root.Left)
	right := balaceTreeHight(root.Right)

	if left == -1 || right == -1 || left - right > 1 || right - left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

