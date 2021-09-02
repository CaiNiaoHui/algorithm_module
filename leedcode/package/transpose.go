package _package

func Transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	matrix1 := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix1[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix1[j][i] = matrix[i][j]
		}
	}

	return matrix1
}
