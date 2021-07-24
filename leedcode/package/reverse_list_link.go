package _package

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {return nil}
	return append(reversePrint(head.Next), head.Val)
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev

		prev = head
		head = temp
	}

	result := make([]int, 0)
	temp := prev

	for temp != nil {
		result = append(result, temp.Val)
		temp = temp.Next
	}

	return result
}


