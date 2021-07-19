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


// deepCopy only copy value
//func DeepCopyList(head *ListNode) *ListNode {
//	newList := &ListNode{
//		Next:  nil,
//		Value: 0,
//	}
//	later := &ListNode{
//		Next:  nil,
//		Value: 0,
//	}
//	//var prev *ListNode
//
//	valList := make([]int, 0)
//	_ = dfs(head, &valList)
//	fmt.Println(valList)
//
//	////
//	//newList.Value = valList[1]
//	//later.Next = newList
//	//newList = later
//	//later.Next = prev // later 需要至空
//	//
//	////
//	//newList.Value = valList[2]
//	//later.Next = newList
//
//
//
//
//	for i := 0; i < len(valList); i++ {
//		newList.Value = valList[i]
//		later.Next = newList
//		copy(later, newList)
//		//newList = later
//		//later.Next = prev // later 需要至空
//		//newList.Value = valList[i] // 获取值
//		//prev = newList		// 把获取的值放到prev下
//		//later.Next = prev	//后面节点把之前节点加入队列
//
//	}
//
//	return later.Next
//}
//
//func dfs(head *ListNode, valList *[]int) *ListNode {
//	if head == nil {
//		return head
//	} else {
//		dfs(head.Next, valList)
//		*valList = append(*valList, head.Value)
//		//fmt.Println(head.Value)
//	}
//	return head
//}



