package _package

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

 

示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20


*/


func MovingCount(m int, n int, k int) int {
	return startdfs(0, 0, m, n, k)
}


func startdfs(i, j, m, n, k int) int {
	visit := make([][]bool, m)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, n)
	}

	return dfs(0, 0, m, n, k, visit)
}

func dfs(i, j, m, n, k int, visit [][]bool) int {
	if i < 0 ||
		j < 0 ||
		i >= m ||
		j >= n ||
		k < thisSum(i, j) ||
		visit[i][j] == true {
		return 0
	}
	visit[i][j] = true

	return 1 + dfs(i-1, j, m, n, k, visit) +
		dfs(i, j-1, m, n, k, visit) +
		dfs(i+1, j, m, n, k, visit) +
		dfs(i, j+1, m, n, k, visit)

}

func thisSum(i, j int) int {
	temp := 0
	for i > 0 && j > 0 {
		temp = temp + i%10 + j%10
		i, j= i/10, j/10
	}
	for i > 0 {
		temp += i%10
		i = i/10
	}
	for j > 0 {
		temp += j%10
		j = j/10
	}
	return temp
}
