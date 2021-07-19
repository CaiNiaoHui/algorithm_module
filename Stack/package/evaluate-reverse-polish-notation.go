package _package

import "strconv"

// tokens 记录传来的参数
// stack 记录结果
func RPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			// string to int
			i, _ := strconv.Atoi(tokens[i])
			stack = append(stack, i)
		}
	}
	return stack[0]
}