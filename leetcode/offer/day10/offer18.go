package day10

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//*删除列表指定元素
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	//*双指针
	pre, cur := head, head.Next

	for ; cur != nil; cur = cur.Next {

		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre = cur //*细节，保存当前节点的前一个指针
	}

	return head
}
