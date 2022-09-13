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

func recursiveNumOccurrences(n *node, target int) int {
	if n == nil {
		return 0
	}
	lastTarget := n.value
	count := recursiveNumOccurrences(n.next, target)
	if target == lastTarget {
		count++
	}
	return count
}

func main() {
	intList := nodeList{}
	intList.append(5)
	intList.append(8)
	intList.append(7)
	intList.append(1)
	intList.append(8)
	intList.append(3)
	intList.append(7)
	intList.append(5)
	intList.append(7)
	intList.append(7)
	intList.append(3)
	intList.append(5)
	intList.append(2)
	intList.append(7)
	intList.append(0)
	toPrint := recursiveNumOccurrences(intList.head, 8)
	fmt.Println(toPrint)
}
