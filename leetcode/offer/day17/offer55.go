package day17

//*二叉树的深度

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//*采用DFS（广度优先）：层序遍历实现
//*使用辅助队列实现
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		tmp := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = append([]*TreeNode{}, tmp...)
		count++
	}
	return count

}

//?使用DFS（深度优先遍历）:递归实现
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

//!平衡二叉树判断：任意节点的左右子树深度不超过1
//?思路：BFS+DFS；得到每个节点；然后计算各个节点子树的深度
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		//*计算左右子树深度
		l := maxDepth(node.Left)
		r := maxDepth(node.Right)
		if l < r { //*防止结果是负数
			l, r = r, l
		}
		if l-r > 1 {
			return false
		}
		//*如果该节点左右子树都存在
		if node.Right != nil && node.Left != nil {
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}

//*DFS+剪枝
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxDepth func(*TreeNode) int

	maxDepth = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		l := maxDepth(tn.Left)
		if l == -1 {
			return -1
		}
		r := maxDepth(tn.Right)
		if r == -1 {
			return -1
		}

		if l < r {
			l, r = r, l
		}
		if l-r > 1 {
			return -1
		}
		return l + 1
	}
	return maxDepth(root) != -1
}
