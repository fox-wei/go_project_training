package day5

import (
	"fmt"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	//*创建二叉树
	root := &TreeNode{Val: 3}
	one := TreeNode{Val: 9}
	two := TreeNode{Val: 20}
	tree := TreeNode{Val: 15}
	four := TreeNode{Val: 7}

	two.Left = &tree
	two.Right = &four
	root.Left = &one
	root.Right = &two

	res := levelOrder(root)
	fmt.Println(res)

	s := []int{0}
	s = s[1:] //*空切片
	fmt.Println(s)

	/*
		queue := NodeQueue{}
		queue.Enqueue(root)
		node := queue.Dequeue()
		fmt.Println(node.Val, len(queue.queue))

		queue.Enqueue(root.Left)
		queue.Enqueue(root.Right)
		fmt.Println(queue.queue)
		node = queue.Dequeue()
		fmt.Println(node, node.Val, len(queue.queue))*/
}

func TesSlice(t *testing.T) {
	s := []int{1}
	s = s[1:]
	fmt.Println(s)
}

func TestBFSOrder(t *testing.T) {
	//*创建二叉树
	root := &TreeNode{Val: 3}
	one := TreeNode{Val: 9}
	two := TreeNode{Val: 20}
	tree := TreeNode{Val: 15}
	four := TreeNode{Val: 7}

	two.Left = &tree
	two.Right = &four
	root.Left = &one
	root.Right = &two

	res := BFSOrder(root)
	fmt.Println(res)
}

func TestBFSOrderR(t *testing.T) {
	//*创建二叉树
	root := &TreeNode{Val: 3}
	one := TreeNode{Val: 9}
	two := TreeNode{Val: 20}
	tree := TreeNode{Val: 15}
	four := TreeNode{Val: 7}

	two.Left = &tree
	two.Right = &four
	root.Left = &one
	root.Right = &two

	res := BFSOrderR(root)
	fmt.Println(res)
}
