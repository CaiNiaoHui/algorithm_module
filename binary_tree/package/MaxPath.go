package _package

type ResultType struct {
	SinglePath int	// 保存单边最大值
	MaxPath int //保存最大值（单边或者两个单边+根的值）
}

func MaxPathSum(root *TreeNode) int {
	length := helper(root)
	return length.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 32),
		}
	}

	// Divide
	left := helper(root.Left)
	right := helper(root.Right)

	// Comquer
	result := ResultType{}
	// 找出最长边
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath + root.Value, 0)
	} else {
		result.SinglePath = max(right.SinglePath + root.Value, 0)
	}
	// 两边的最大值
	maxpath := max(left.SinglePath, right.SinglePath)
	// 两边加根最大值
	result.MaxPath = max(left.SinglePath + right.SinglePath + root.Value, maxpath)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}



