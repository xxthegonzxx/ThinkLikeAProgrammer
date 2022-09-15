package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func recursiveInsertBinarySearch(n *treeNode, num int) int {
	if n == nil {
		return 0
	}
	if n.data == num {
		fmt.Printf("%d, Duplicate Value already in tree...skipping.\n", n.data)
	}
	if n.data > num {
		if n.left == nil {
			n.left = &treeNode{
				data: num,
			}
			return 0
		}
		return recursiveInsertBinarySearch(n.left, num)
	}
	if n.data < num {
		if n.right == nil {
			n.right = &treeNode{
				data: num,
			}
			return 0
		}
		return recursiveInsertBinarySearch(n.right, num)
	}
	return 0
}

func recursiveBinarySum(t *treeNode) (int, int) {
	if t == nil {
		return 0, 0
	}
	leftInts, leftCount := recursiveBinarySum(t.left)
	rightInts, rightCount := recursiveBinarySum(t.right)
	nodeCount := leftCount + rightCount + 1
	sum := t.data + leftInts + rightInts

	return sum, nodeCount
}

func average(n *treeNode) int {
	sum, nodeCount := recursiveBinarySum(n)
	return sum / nodeCount
}

func main() {

	rootNode := &treeNode{data: 10}
	recursiveInsertBinarySearch(rootNode, 5)
	recursiveInsertBinarySearch(rootNode, 15)

	average := average(rootNode)
	fmt.Println(average)

}
