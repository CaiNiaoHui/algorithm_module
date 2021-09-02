package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Golang regular expressions example"
	// 判断是否是Golang开头
	match, err := regexp.MatchString(`^Golang`, str)
	fmt.Println("Match: ", match, " Error: ", err)
}
