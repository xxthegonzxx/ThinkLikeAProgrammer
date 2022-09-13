package main

import "fmt"

type node struct {
	value int
	next  *node
}

type nodeList struct {
	head *node
}

func (l *nodeList) append(num int) {
	newNode := node{
		value: num,
		next:  l.head,
	}
	l.head = &newNode
}

func recursiveOddParity(n *node) bool {
	if n == nil {
		return false
	}
	traverse := recursiveOddParity(n.next)
	if traverse {
		return n.value != 1
	}
	return n.value == 1
}

func main() {
	intList := nodeList{}
	intList.append(1)
	intList.append(0)
	intList.append(1)
	intList.append(1)
	intList.append(0)
	toPrint := recursiveOddParity(intList.head)
	fmt.Println(toPrint)
}
