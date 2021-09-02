package _package

import (
	"encoding/json"
)


func Normal_DeepCopy(list []int) []int {
	result := make([]int, len(list))
	copy(list, result)

	return result
}

func Normal_DeepCopy2(s01 []string) []string {
	s02 := make([]string, len(s01))
	data, _ := json.Marshal(s01)
	json.Unmarshal(data, &s02)
	
	return s02
}

func Deep_link_my(head *ListNode) *ListNode {
	head2 := &ListNode{
		Next:  nil,
		Value: 0,
	}
	DumpyHead2 := head2
	DumpyHead1 := head

	if DumpyHead1 == nil {
		return DumpyHead2.Next
	}
	// 拷贝
	copyL(DumpyHead1, DumpyHead2)

	return head2.Next
}

func copyL(l1 *ListNode, l2 *ListNode)  {
	if l1 == nil {
		return
	}
	l2.Next = &ListNode{
		Value: l1.Value,
		Next:  nil,
	}
	copyL(l1.Next, l2.Next)
}



