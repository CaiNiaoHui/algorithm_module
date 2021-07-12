package _package

// 这里的head节点可能是链表中的任意一个节点
// 删除重复元素，使每个元素只出现一次
func DeleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {

		for current.Next != nil && current.Value == current.Next.Value {
			current.Next = current.Next.Next
		}

		current = current.Next
	}

	return head
}


// 使用 dummy node 辅助删除
// 删除所有含有的重复节点
/*func DeleteDuplicatesAll(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Value: 0}
	dummy.Next = head
	head = dummy

	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Value == head.Next.Next.Value {
			// 记录已经删除的值
			rmVal = head.Next.Value
			for head.Next != nil && head.Next.Value == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
*/

//删除所有重复节点
func DeleteDuplicatesAll(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 辅助节点防止头节点被删除
	dummy := &ListNode{Value: 0}
	dummy.Next = head
	head = dummy
	// 记录删除节点的值
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Value == head.Next.Next.Value {
			rmVal = head.Next.Value
			for head.Next != nil && rmVal == head.Next.Value{
				head.Next = head.Next.Next
			}
		}else {
			head = head.Next
		}
	}
	return dummy.Next
}


