package day19

//*重建二叉树：根据前序遍历和中序遍历还原二叉树
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	inorderIdx := 0
	//*确定根节点在inorder位置
	for i, val := range inorder {
		if rootVal == val {
			inorderIdx = i
			break
		}
	}

	//*根据inodrderIdx确定中序遍历的左右子树
	leftInorder := inorder[:inorderIdx]    //*中序遍历左子树
	rightInorder := inorder[inorderIdx+1:] //*中序遍历右子树

	//*根据leftInoder的长度确定前序遍历的左右子树
	leftPreorder := preorder[1 : len(leftInorder)+1] //*前序遍历的左子树
	rightPreorder := preorder[len(leftInorder)+1:]   //*前序遍历的右子树，小细节问题

	//*递归遍历
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
