package main

import (
	"fmt"
	"regexp"
)

/*
Compile() 或者 MustCompile()创建一个编译好的正则表达式对象。
假如正则表达式非法，那么Compile()方法回返回error,
而MustCompile()编译非法正则表达式时不会返回error，而是回panic。
如果你想要很好的性能，不要在使用的时候才调用Compile()临时进行编译，
而是预先调用Compile()编译好正则表达式对象：

regexp1, err := regexp.Compile(<code>regexp</code>)
regexp2 := regexp.MustCompile(<code>regexp</code>)
*/

func main() {
	p := "a*"
	s := "aaa"
	regexp1 := regexp.MustCompile("^" + p + "$")
	regexp2 := regexp1.MatchString(s)
	fmt.Println(regexp1 , regexp2)

}
