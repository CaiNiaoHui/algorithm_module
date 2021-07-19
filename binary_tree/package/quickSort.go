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

// 快排思想: 以最后一个节点为标记点,
// 从第一个开始遍历数组, 比标记点小的就合第i交换i++，
// 最后遍历完在交换i和标记
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

