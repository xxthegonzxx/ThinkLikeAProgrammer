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

func (ptr *treePtr) append(num int) {
	newNode := treeNode{
		data:  num,
		left:  ptr.head,
		right: ptr.head,
	}
	ptr.head = &newNode
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

func main() {
	intBinaryList := treePtr{}
	intBinaryList.append(10)
	intBinaryList.append(20)
	intBinaryList.append(30)
	intBinaryList.append(40)
	intBinaryList.append(45)
	intBinaryList.append(50)
	intBinaryList.append(60)
	intBinaryList.append(70)
	intBinaryList.append(100)
	isBinaryHeap := recurseIsBinaryHeap(intBinaryList.head)
	fmt.Println(isBinaryHeap)
}
