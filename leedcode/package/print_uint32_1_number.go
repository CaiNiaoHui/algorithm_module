package _package

func Search1Numbers(s uint32) int {

	count := 0

	for s > 0 {
		s = s & (s-1)
		count++
	}
	return count
}
