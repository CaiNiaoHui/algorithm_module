package main

import (
	"fmt"
	_package "github.com/CaiNiaoHui/algorithm_module/leedcode/package"
)

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

 

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

*/

func main() {
	//fmt.Println("print: fib list: ", _package.Fib(45))
	//
	//a := []int {
	//	2, 3, 1, 0, 2, 5, 3,
	//}
	//fmt.Println("print: repeat list: ", _package.FindRepeatNumber2(a))
	//
	//b := [][]int {
	//	{1,   4,  7, 11, 15},
	//	{2,   5,  8, 12, 19},
	//	{3,   6,  9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}
	//fmt.Println("print: find target value: ", _package.FindNumberIn2DArray(b, 20))
	//
	//fmt.Println("sport chance numbers: ", _package.MovingCount(3, 1, 0))
	//
	//board := [][]byte{
	//	{'A','B','C','E'},
	//	{'S','F','C','S'},
	//	{'A','D','E','E'},
	//}
	//word := "ABCCED"
	//fmt.Println("chart is : ", _package.Exist(board, word))
	//
	//fmt.Println("裁剪绳子n段相乘得到的最大值为: ", _package.CuttingRope1(100))
	//
	//nums := []int{
	//	-1,0,3,5,9,12,
	//}
	//fmt.Println("二分查找对应数字的下标: ", _package.Search(nums, 9))

	fmt.Println(_package.IsNumber("5e2"))

	fmt.Println(_package.Search1Numbers(11))

	matrix := [][] int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(_package.SpiralOrder(matrix))

	fmt.Println(_package.MyPow(2.1, 3))

	fmt.Println(_package.Sumzero(5))


	fmt.Println(_package.SubarraySum([]int{1,1,1}, 2))

	fmt.Println(_package.FindKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}





