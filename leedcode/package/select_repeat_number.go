package _package

// 在数组中挑选中随机一个重复数

func FindRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return num
		}
		m[num] = true
	}
	return -1
}

func FindRepeatNumber2(nums []int) int {

	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if map1[nums[i]] == 0{
			map1[nums[i]] = 1
		} else {
			return nums[i]
		}
	}

	return -1
}