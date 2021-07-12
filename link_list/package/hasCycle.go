package _package

func HasCycle(head *ListNode) bool {
	// 快慢指针
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {

		if slow == fast {
			return  true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}

// 找到第入环点
func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 快指针置头指针, 同等步伐直到相遇
			fast = head
			slow = slow.Next // 注意
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
		// 快慢指针走自己的步伐
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}