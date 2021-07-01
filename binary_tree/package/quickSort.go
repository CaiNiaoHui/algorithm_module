package _package

// 快速排序 属于divide分治法，但是没有归并

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	// divide
	if start < end {
		mid := partition(nums, start, end)
		quickSort(nums, 0, mid-1)
		quickSort(nums, mid+1, end)
	}
}

func partition(nums []int, start, end int) int {
	i, p := start, nums[end]
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

