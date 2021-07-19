package main

import (
	"fmt"
	_package "github.com/CaiNiaoHui/algorithm_module/Stack/package"
)

func main() {
	// 初始化一个栈
	minStack := _package.Construction()
	minStack.Push(2)
	minStack.Push(4)
	minStack.Push(1)
	minStack.PrintStack()

	token := []string{"2", "1", "+", "3", "*"}
	fmt.Println("RPN结果：", _package.RPN(token))

	s := "2[abc]3[cd]ef"
	fmt.Println(_package.DecodeString1(s))


}
