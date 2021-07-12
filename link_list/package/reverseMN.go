package _package

//func ReverseBetweenMN(head *ListNode, m int, n int) *ListNode {
//	// 先遍历到m处，反转，再拼接后续
//	if head == nil {
//		return head
//	}
//	// 头部变化使用dummy node
//	dummy := &ListNode{Value: 0}
//	dummy.Next = head
//	head = dummy
//
//	// 先遍历到m处
//	var pre *ListNode
//	var i int
//	for i = 0; i < m; i++ {
//		pre = head
//		head = head.Next
//	}
//	// 遍历之后
//	j := i
//	var next *ListNode
//	// 中间节点连接
//	var mid = head
//	for head != nil && j <= n {
//		// 第一次循环
//		temp := head.Next
//		head.Next = next
//		next = head
//		head = temp
//
//		j++
//	}
//	// 循环结束
//	pre.Next = next		// next是反转之后的头节点
//	mid.Next = head		// 后面队列链接给mid
//
//	return dummy.Next
//}

func ReverseMN(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	// 头节点改变
	dummy := &ListNode{Value: 0}
	dummy.Next = head
	head = dummy

	// 遍历到m
	var i int
	var prev1 *ListNode
	for i = 0; i < m; i++ {
		prev1 = head
		head = head.Next
	}

	// 反转部分
	j := i
	mid := head  // 拼接中间
	var prev *ListNode
	for j <= n && head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp

		j++
	}

	// 拼接
	prev1.Next = prev
	mid.Next = head


	return dummy.Next
}