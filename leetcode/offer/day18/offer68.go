package day18

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//!查找二叉搜索树的公共祖先节点-特例
/*
*祖先节点：节点p是root的子树，则root是p的祖先节点
*最近公共祖先节点：root是p, q的某公共祖先节点，若root的左右子树不是p,q的公共祖先节点，则root是p,q的最近公共祖先节点
 */
//*二叉搜索树的左子树小于右子树
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p //*保证p<q
	}

	for root != nil {
		if p.Val > root.Val {
			root = root.Right
		} else if q.Val < root.Val {
			root = root.Left
		} else {
			break
		}
	}

	return root
}

//&递归实现
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	}
	return root
}

//!查找二叉树的公共祖先节点-先序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
