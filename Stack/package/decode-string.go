package _package

import "strconv"

/*
给定一个经过编码的字符串，返回它解码后的字符串。
s = "3[a]2[bc]", 返回 "aaabcbc". s = "3[a2[c]]",
返回 "accaccacc". s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/
// 1. stack存储结果， 临时存储除']'的值
// 2. temp存储需要重复的字符串
// 3. 使用idx来处理num的位数

func DecodeString(s string) string {

	if len(s) == 0 {
		return ""
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			// pop '['
			stack = stack[:len(stack)-1]
			// pop num
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			// 注意索引边界
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				// 把字符正向放回到栈里面
				for j := len(temp) - 1; j >= 0; j-- {
					stack = append(stack, temp[j])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func DecodeString1(s string) string {
	if len(s) == 0 {
		return ""
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				temp = append(temp, v)
				stack = stack[:len(stack)-1]
			}
			// '['
			stack = stack[:len(stack)-1]
			// num
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '1' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			// 取出索引边界，可能是两位是
			num := stack[len(stack)+1-idx:]
			stack = stack[:len(stack)+1-idx]
			count, _ := strconv.Atoi(string(num))	 // 先转换到string在转换到int
			for j := 0; j < count; j++ {
				for t := len(temp)-1; t >= 0; t-- {
					stack = append(stack, temp[t])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}