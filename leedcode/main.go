package main

import (
	"fmt"
	_package "github.com/CaiNiaoHui/algorithm_module/leedcode/package"
)

func main() {
	fmt.Println("print: fib list: ", _package.Fib(45))

	a := []int {
		2, 3, 1, 0, 2, 5, 3,
	}
	fmt.Println("print: repeat list: ", _package.FindRepeatNumber2(a))

	b := [][]int {
		{1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println("print: find target value: ", _package.FindNumberIn2DArray(b, 20))

	fmt.Println("sport chance numbers: ", _package.MovingCount(3, 1, 0))

	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	word := "ABCCED"
	fmt.Println("chart is : ", _package.Exist(board, word))

	fmt.Println("裁剪绳子n段相乘得到的最大值为: ", _package.CuttingRope1(100))

}





