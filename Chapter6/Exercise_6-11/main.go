package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func recurseInsertBinarySearchTree(n *treeNode, num int) error {
	if n == nil {
		return fmt.Errorf("Tree is NIL.")
	}
	if n.data > num {
		if n.left == nil {
			n.left = &treeNode{
				data: num,
			}
			return nil
		}
		return recurseInsertBinarySearchTree(n.left, num)
	}
	if n.data < num {
		if n.right == nil {
			n.right = &treeNode{
				data: num,
			}
			return nil
		}
		return recurseInsertBinarySearchTree(n.right, num)

	}
	return nil
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
	// Example Binary Search Tree
	//	 	100
	//	   /   \
	//	 70     800
	//	/  \    /  \
	// 50  90  600 4000
	rootNode := &treeNode{data: 100}

	recurseInsertBinarySearchTree(rootNode, 90)
	recurseInsertBinarySearchTree(rootNode, 800)
	recurseInsertBinarySearchTree(rootNode, 70)
	recurseInsertBinarySearchTree(rootNode, 600)
	recurseInsertBinarySearchTree(rootNode, 50)
	recurseInsertBinarySearchTree(rootNode, 4000)

	rootNode.PrintInorder()
}
