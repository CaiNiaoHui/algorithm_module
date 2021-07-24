package _package

import "sort"

// 最快寻找有序矩阵中目标数

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	for _,nums:=range matrix{
		i:=sort.SearchInts(nums,target)
		if i<len(nums)&&target==nums[i]{
			return true
		}
	}
	return false
}

