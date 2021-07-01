package _package

// 层序遍历
func Bfs_Traverse(root *TreeNode) [][]int {
	result := bfs(root)
	return result
}

func bfs(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	// Comquer
	querer := make([]*TreeNode, 0)
	querer = append(querer, root)

	for len(querer) > 0 {
		list := make([]int, 0)
		l := len(querer)

		for i := 0; i < l; i++ {
			level := querer[0]
			querer = querer[1:]
			list = append(list, level.Value)

			if level.Left != nil {
				querer = append(querer, level.Left)
			}
			if level.Right != nil {
				querer = append(querer, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

// 从下到上层序遍历

func Bfs_Traverse_reverse(root *TreeNode) [][]int {
	result := bfs(root)
	reverse(result)
	return result
}

func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// bfs层序Z字形遍历

func Bfs_Traverse_Z(root *TreeNode) [][]int {
	result := bfs_Z(root)
	return result
}

func bfs_Z(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	// Comquer
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag := false

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)

		for i := 0; i < l; i++ {
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Value)

			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		if flag == true {
			reverse_list(list)
		}
		flag = !flag
		result = append(result, list)
	}
	return result
}

func reverse_list(nums []int) {
	for i, j := 0, len(nums)-1; i<j ; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}