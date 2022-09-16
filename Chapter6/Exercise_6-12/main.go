package main

import (
	"fmt"
	"sort"
)

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

func recursiveBinaryMedian(t *treeNode) int {
	_, nodeCount := recursiveBinarySum(t)
	findEvenMiddle := (nodeCount / 2)
	findOddMiddle := (nodeCount + 1) / 2

	intSlice := inOrderTraversalList(t)
	sort.Ints(intSlice)
	if t == nil {
		return 0
	}
	if nodeCount%2 == 0 {
		firstMiddle := intSlice[findEvenMiddle-1]
		secondMiddle := intSlice[findEvenMiddle]
		evenMedian := firstMiddle + secondMiddle/2
		return evenMedian
	} else {
		oddMedian := intSlice[findOddMiddle-1]
		return oddMedian
	}
}

func inOrderTraversalList(t *treeNode) []int {
	intSlice := []int{t.data}
	if t == nil {
		return intSlice
	}
	if t.left != nil {
		intSlice = append(intSlice, inOrderTraversalList(t.left)...)
	}

	if t.right != nil {
		intSlice = append(intSlice, inOrderTraversalList(t.right)...)
	}
	return intSlice
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

func average(t *treeNode) int {
	sum, nodeCount := recursiveBinarySum(t)
	return sum / nodeCount
}

func mode(t *treeNode, target int) int {
	if t == nil {
		return 0
	}
	if t.data == target {
		count := mode(t.left, target) + mode(t.right, target) + 1
		return count
	}
	count := mode(t.left, target) + mode(t.right, target)
	return count
}

func main() {

	rootNode := &treeNode{data: 10}
	recursiveInsertBinarySearch(rootNode, 5)
	recursiveInsertBinarySearch(rootNode, 15)
	recursiveInsertBinarySearch(rootNode, 20)
	recursiveInsertBinarySearch(rootNode, 1)

	average := average(rootNode)
	fmt.Println("Average:", average)

	median := recursiveBinaryMedian(rootNode)
	fmt.Println("Median:", median)
	mode := mode(rootNode, 5)
	fmt.Println(mode)
}
