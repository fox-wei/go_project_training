package main

//*链表反转，以数组形式逆序返回节点的值

type ListNode struct {
	Val  int
	Next *ListNode
}

//*链表迭代
func reversePrint(head *ListNode) []int {
	tem := []int{}
	if head == nil {
		// tem = append(tem, head.Val)
		return []int{}
	}

	for next := head; next.Next != nil; next = next.Next {
		tem = append(tem, next.Val)
	}

	l := []int{}
	for i := len(tem); i >= 0; i-- {
		l = append(l, tem[i])
	}

	return l
}

//?反转链表
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	var pre *ListNode = nil
	cur := head

	//?反转 offer24
	for cur != nil {
		tem := cur.Next
		cur.Next = pre
		pre = cur
		cur = tem
	}

	//*迭代
	tem := []int{}
	for next := pre; next != nil; next = next.Next {
		tem = append(tem, next.Val)
	}
	return tem
}
