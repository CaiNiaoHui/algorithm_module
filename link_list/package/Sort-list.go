package _package

func SortList(head *ListNode) *ListNode {
	// 归并排序 O(nlogn)
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	// head 为 nil返回
	if head == nil || head.Next == nil {
		return head
	}
	// 找到 中间节点
	middle := findMiddle(head)
	// 切割节点
	tail := middle.Next
	middle.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)
	// 归并有序链表
	result := mergeTwoList(left, right)

	return result
}

func findMiddle(head *ListNode) *ListNode {
	// 快慢指针
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{Value: 0}
	dummy := result

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			result.Next = l1
			l1 = l1.Next
			result = result.Next
		} else {
			result.Next = l2
			l2 = l2.Next
			result = result.Next
		}
	}

	for l1 != nil {
		result.Next = l1
		result = result.Next
		l1 = l1.Next
	}
	for l2 != nil {
		result.Next = l2
		result = result.Next
		l2 = l2.Next
	}

	return dummy.Next
}




