package day14

//*二叉搜索树的第k个最大值
//?二叉搜索树的特点：左节点小于根节点，右节点大于根节点；
//*使用中序遍历可构成有序数据；这里使用中序：右、根、左

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	res := -1

	//*中序遍历
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)

	return res
}
