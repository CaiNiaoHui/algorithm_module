package _package

// 分治法
func Dfs_downToUp(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	// 递归
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	// 归并
	result = append(result,root.Value)
	result = append(result, left...)
	result = append(result, right...)

	return result
}


/// 分治法模版
//
//type ResultType []int
//
//func traversal(root *TreeNode) ResultType {
//	//nil
//	result := make(ResultType, 0)
//	if root == nil {
//		return  nil
//	}
//	// Divide
//	ResultType left := traversal(root.Left)
//	ResultType right := traversal(root.Right)
//
//	//Conquer
//	ResultType result = Merge from left and right
//
//	return result
//}

