package day11

//*找出两个链表的相交节点,常规双循环遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA

L:
	for p1 != nil {
		p2 := headB
		for p2 != nil {
			if p1 == p2 {
				break L
			}
			p2 = p2.Next
		}

		p1 = p1.Next
	}

	if p1 == nil {
		return nil
	}

	return p1
}

//*使用双指针，p1从A开始到B遍历，p2从B开始到A遍历
//*两者相等即公共点
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB

	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
