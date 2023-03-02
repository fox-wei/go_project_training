package day6

//*二叉树镜像
//** Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//?递归实现；暂存左子树，右子树作为左子树的返回值
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)

	return root
}

//*非递归实现：使用队列
func mirrorTreeByStack(root *TreeNode) *TreeNode {
	if root != nil {
		return nil
	}
	queue := []*TreeNode{root} //*辅助队列

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}

	return root
}
