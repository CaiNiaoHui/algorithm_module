package _package

// 切分链表，使所有大于x的值都在后面
func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	// 调整头节点
	dummyHead := &ListNode{Value: 0}
	dummyTail := &ListNode{Value: 0}
	dummyHead.Next = head
	head = dummyHead
	tail := dummyTail

	// 遍历
	for head.Next != nil {
		// 小于 移位
		if head.Next.Value < x {
			head = head.Next
		} else {
		// 大于 剔除 存到另一个链表中
			t := head.Next
			head.Next =head.Next.Next
			tail.Next = t
			tail = tail.Next
		}
	}

	// 合并
	tail.Next = nil	// 置空
	head.Next = dummyTail.Next	// 拼接

	return dummyHead.Next
}