package _package

func CopyRandomList(head *ListNodeR) *ListNodeR {
	if head == nil {
		return head
	}
	// 复制节点
	// 1 > 2 > 3
	cur := head
	for cur != nil {
		clone := &ListNodeR{Value: cur.Value, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	// 处理random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离两个链表
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	// 原始链表头： head 1 > 2 > 3
	// 克隆链表头: cloneHead 1' > 2' > 3'
	return cloneHead
}

