package _package

import "math"

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
提示：

2 <= n <= 58

*/

// 小数
func CuttingRope1(n int) int {
	if n <= 3 {
		return n-1
	}
	a, b := n/3, n%3

	var result float64
	if b == 0 {
		result = math.Pow(3, float64(a))
	}else if b == 1{
		result = math.Pow(3, float64(a-1))*4
	}else {
		result = math.Pow(3, float64(a))*2
	}

	return int(result)
}

// 大数 取模
func CuttingRope2(n int) int {
	const mod = 1e9+7

	if n <= 3 {
		return n-1
	}
	a, b := n/3, n%3
	result := 1
	if b == 0 {
		for i := 1; i <= a; i++ {
			result = result * 3 % mod
		}
	} else if b ==1 {
		for i := 1; i <= a-1; i++ {
			result = result * 3 % mod
		}
		result = result*4 % mod
	} else {
		for i := 1; i <= a; i++ {
			result = result * 3 % mod
		}
		result = result*2 % mod
	}
	return result

}

