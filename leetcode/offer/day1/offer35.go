package main

//*普通列表复制

func copyListNode(head *ListNode) *ListNode {

	pre := new(ListNode)
	pre.Val = head.Val
	pre.Next = nil
	//?遍历
	for next := head.Next; next != nil; next = next.Next {
		//?生成新节点
		node := new(ListNode)
		node.Val = next.Val
		node.Next = next.Next

		pre.Next = node
		pre = node //*缓存本节点
	}
	return pre
}

//!复杂链表的复制

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	//?构建map，复制节点,建立原节点->新节点的映射
	listMap := make(map[*Node]*Node)
	for cur != nil {
		listMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	//&构建新表的next和random指向
	for cur != nil {
		listMap[cur].Next = listMap[cur.Next]
		listMap[cur].Random = listMap[cur.Random]
		cur = cur.Next
	}

	return listMap[head]
}
