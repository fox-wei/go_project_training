package day14

//*二叉搜索树转换成双向循环链表
//*使用中序遍历，pre,head指针修改指针的前后关系

func Convert(pRootOfTree *TreeNode) *TreeNode {

	// write code here
	var pre, head *TreeNode
	var dfs func(*TreeNode)
	//*中序遍历
	dfs = func(cur *TreeNode) {
		//*中序遍历
		if cur == nil {
			return
		}
		dfs(cur.Left)
		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur
		dfs(cur.Right)
	}
	dfs(pRootOfTree)
	pre.Right = head
	head.Left = pre
	return head
}
