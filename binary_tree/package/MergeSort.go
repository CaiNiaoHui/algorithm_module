package _package

func MergeSort(array []int) []int {
	return mergeSort(array)
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	// divide
	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	// merge
	result := merge(left, right)

	return result
}

func merge(left, right []int) []int {
	l, r := 0, 0
	list := make([]int, 0)

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			list = append(list, left[l])
			l++
		} else {
			list = append(list, right[r])
			r++
		}
	}
	// 归并剩下的
	list = append(list, left[l:]...)
	list = append(list, right[r:]...)

	return list
}