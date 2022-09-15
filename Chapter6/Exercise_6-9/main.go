package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

type treePtr struct {
	head *treeNode
}

func recurseIsBinaryHeap(n *treeNode) bool {
	if n == nil {
		return false
	}
	if n.left == nil && n.right == nil {
		return false
	}
	recurseIsBinaryHeap(n.left)
	recurseIsBinaryHeap(n.right)
	if n.data > n.left.data && n.data > n.right.data {
		return true
	}
	return false
}

// Inorder traversal is a way of traversing through the binary tree
// by where the left child is visited first followed by the parent
// and then the right child of the node.
func (t *treeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.left.PrintInorder()
	fmt.Println(t.data)
	t.right.PrintInorder()
}

func main() {
	// Example Binary Tree
	//	 	100
	//	   /   \
	//	 90     80
	//	/  \   /  \
	// 70  60 50  40

	node7 := &treeNode{40, nil, nil}
	node6 := &treeNode{50, nil, nil}
	node5 := &treeNode{60, nil, nil}
	node4 := &treeNode{70, nil, nil}
	node3 := &treeNode{80, node6, node7}
	node2 := &treeNode{90, node4, node5}
	rootNode := &treeNode{100, node2, node3}

	isBinaryHeap := recurseIsBinaryHeap(rootNode)
	fmt.Println("Is Binary Heap?", isBinaryHeap)
	rootNode.PrintInorder()
}
