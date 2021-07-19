package _package

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next	// 保存head的下一个节点和之后的队列
		head.Next = prev	// 当前取到的节点为prev的前节点

		prev = head		// 移动prev 得到结果的head队列赋值给prev
		head = temp		// 移动head
	}

	return prev
}
