package day19

import "math"

//&二叉搜索树的性质：根节点的值大于左节点；小于右节点；
//&中序遍历结果是递增
//*根据后序遍历结果判断是否为二叉搜索树（值不相同）
//*思路：最后一个节点是根节点；小于根节点的为左子树的部分；大于为右子树；依次递归运算
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	l := len(postorder)
	root := postorder[l-1] //*根节点

	//*确定左子树部分
	i := 0
	for ; i < l-1; i++ {
		if postorder[i] > root { //*达到右子树
			break
		}
	}

	j := i //*右子树开始
	for ; j < l-1; j++ {
		if postorder[j] < root { //*右子树部分一定大于根节点
			return false
		}
	}

	return verifyPostorder(postorder[:i]) && verifyPostorder(postorder[i:l-1])
}

//^非递归实现：利用二叉树搜索树的性质：右子树一定大于左子树
func verifyPostorder2(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	stack := []int{}
	root := math.MaxInt

	//*必须从最后一个开始，否则某些用例无法通过：1， 14， 13， 2返回true
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false //*左边一定小于右边
		}
		for len(stack) > 0 && postorder[i] < stack[len(stack)-1] {
			root = stack[len(stack)-1] //*保存上一个节点
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
