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

func recursiveSumNumbers(n *node) int {
	if n == nil {
		return 0
	}
	if n.next == nil {
		return n.value
	}
	sum := recursiveSumNumbers(n.next) + n.value
	return sum
}

func main() {
	intList := nodeList{}
	intList.append(5)
	intList.append(5)
	toPrint := recursiveSumNumbers(intList.head)
	fmt.Println(toPrint)
}
