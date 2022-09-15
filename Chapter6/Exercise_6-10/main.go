package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func recurseBinarySearchTree(n *treeNode) bool {
	if n == nil {
		return false
	}
	if n.left == nil && n.right == nil {
		return false
	}
	recurseBinarySearchTree(n.left)
	recurseBinarySearchTree(n.right)
	if n.data > n.left.data && n.data < n.right.data {
		return true
	}
	return false
}

func main() {
	node7 := &treeNode{4000, nil, nil}
	node6 := &treeNode{50, nil, nil}
	node5 := &treeNode{600, nil, nil}
	node4 := &treeNode{70, nil, nil}
	node3 := &treeNode{800, node6, node7}
	node2 := &treeNode{90, node4, node5}
	rootNode := &treeNode{100, node2, node3}

	toPrint := recurseBinarySearchTree(rootNode)
	fmt.Println(toPrint)
}
