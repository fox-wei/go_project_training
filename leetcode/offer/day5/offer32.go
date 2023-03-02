package day5

//*二叉树遍历

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//*使用队列保存节点，在获取节点的值
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{} //*保存结果
	queue := NodeQueue{}
	queue.Enqueue(root)

	for queue.queue != nil {
		//*出队
		node := queue.Dequeue()
		// fmt.Println(node, len(queue.queue))
		// if node == nil {
		// 	break
		// }
		res = append(res, node.Val)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}

		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
	return res
}

type NodeQueue struct {
	queue []*TreeNode
}

func (n *NodeQueue) Enqueue(node *TreeNode) {
	n.queue = append(n.queue, node)
}

func (n *NodeQueue) Dequeue() *TreeNode {
	if len(n.queue) == 0 {
		return nil
	}
	tmp := n.queue[0]
	if len(n.queue) >= 2 {
		n.queue = append([]*TreeNode{}, n.queue[1:]...)
	} else {
		n.queue = nil
	}
	return tmp
}

//!简化版本，slice当作队列使用
func Order(root *TreeNode) []int {
	res := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}

//!II同层遍历[[1], [2, 4]]
func BFSOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue) //*当前层的节点数量
		tmp := []int{}

		for i := l; i > 0; i-- {
			node := queue[0]
			tmp = append(tmp, node.Val)
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}

//^同层遍历，单：从左到右，双则相反
func BFSOrderR(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue) //*当前层的节点数量
		tmp := make([]int, l)

		for i := l; i > 0; i-- {
			node := queue[0]
			if (len(res)+1)%2 == 0 { //*偶数层
				tmp[i-1] = node.Val
			} else {
				tmp = append(tmp, node.Val)
			}
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}
