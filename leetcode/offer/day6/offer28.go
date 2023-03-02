package day6

//*判断二叉树是否对称
//?二叉树对称：root.L.l==root.R.r；root.L.r==root.R.l
//*即左子树的左节点等于右子树的右节点；右节点等于左节点

//^使用两个指针进行递归实现
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

//*递归函数
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
