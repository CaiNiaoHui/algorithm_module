package _package

// 斐波拉契数列， 注意大数相加需要取模

func Fib(n int) int {
	const mod = 1e9+7
	f0, f1 := 0, 1
	if n == 0 {
		return f0
	}
	if n== 1 {
		return f1
	}
	for i := 2; i <= n; i++ {
		f2 := (f0 + f1) % mod
		f0 , f1 = f1, f2
	}
	return f1
}
