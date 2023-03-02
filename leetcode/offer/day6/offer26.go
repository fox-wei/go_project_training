package day6

//*判断二叉树B是否为A的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (checkEquel(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func checkEquel(A, B *TreeNode) bool {
	if B == nil { //?遍历结束，匹配成功
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return checkEquel(A.Left, B.Left) && checkEquel(A.Right, B.Right)
}
