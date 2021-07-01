package _package

func LowerCommondRoot(root, l, r *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 就是root节点 返回root节点
	if l == root || r == root {
		return root
	}

	// Divide
	left := LowerCommondRoot(root.Left, l, r)
	right := LowerCommondRoot(root.Right, l, r)

	// Comquer
	// 两边都不为空，根节点为祖先
	if left != nil || right != nil {
		return root
	}

	if left != nil {
		return root
	}

	if right != nil {
		return root
	}

	return nil
}
