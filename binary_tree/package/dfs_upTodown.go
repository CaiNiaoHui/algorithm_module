package _package


func Dfs_upTodown(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Value)
	dfs(root.Left, result)
	dfs(root.Right, result)
}



