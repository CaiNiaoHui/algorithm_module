package _package

// 1. find middle
// 2. reverse tail link list
// 3. compare value
func IsPalidrome(head *ListNode) bool {
	// 1 2 1 nil
	if head == nil {
		return true
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 需要反转的节点
	result := slow.Next
	tail := reverseList(result)
	// 断开链表
	slow.Next = nil

	for head != nil && tail != nil {
		if head.Value != tail.Value {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var prev *ListNode
	// 交换
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}

	return prev

}
