package _package

// 合并升序链表
func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {

	// 改变头节点
	dummy := &ListNode{Value: 0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}

	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head.Next = head
	}

	return dummy.Next
}
