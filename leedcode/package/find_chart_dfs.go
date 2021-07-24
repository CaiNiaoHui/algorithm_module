package _package

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
 

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

*/



func Exist(board [][]byte, word string) bool {

	return dfs_find(board, word)

}

func dfs_find(board [][]byte, word string) bool {
	k := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs_chart(board, word, i, j, k) {
				return true
			}
		}
	}
	return false
}

func dfs_chart(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
		return  false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	result := dfs_chart(board, word, i+1, j, k+1) ||
		dfs_chart(board, word, i-1, j, k+1) ||
		dfs_chart(board, word, i, j+1, k+1) ||
		dfs_chart(board, word, i, j-1, k+1)
	board[i][j] = word[k]

	return result
}

