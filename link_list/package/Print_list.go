package _package

import "fmt"

func Print_list(head *ListNode) {
	current := head
	flag := false
	fmt.Print("遍历链表: ")
	for current != nil {
		if flag == true {
			fmt.Print(" -> ")
		}
		fmt.Printf("%d", current.Value)
		current = current.Next
		flag = true
	}
	fmt.Println()
}



