package _package

/*
青蛙跳台

一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

 */

func numWays(n int) int {
	const mod = 1e9 + 7
	t0 , t1 := 1, 1
	for i := 2; i <= n; i++ {
		result := (t0 + t1) % mod
		t0 , t1 = t1, result
	}
	return t1
}

