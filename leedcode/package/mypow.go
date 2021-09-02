package _package

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1/x
		n = -n
	}
	result := 1.0
	for n > 1 {
		//	说明是奇数
		if n & 1 == 1 {
			result *= x
			n = n-1
		}else {
			x *= x
			n = n >> 1
		}
	}
	result *= x
	return result

}
