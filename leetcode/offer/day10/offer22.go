package day10

//*返回列表倒数第K节点
//?使用双指针，former，latter，former先前进k步；
//?然后两个指针共同前进，直到former==nil，找到倒数第k个节点

func getKthFromEnd(head *ListNode, k int) *ListNode {
	l, f := head, head //?双指针

	//?f前进K步
	for i := 0; i < k; i++ {
		f = f.Next
	}

	for f != nil {
		f = f.Next
		l = l.Next
	}

	return l
}
