package _package

// 中序遍历，判断数组是否有序
func Is_BST(root *TreeNode) bool {
	result := make([]int, 0)
	Inorder(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func Inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	// divide
	Inorder(root.Left, result)
	*result = append(*result, root.Value)
	Inorder(root.Right, result)

}
