package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type stringCollection struct {
	head   *node
	length int
}

// addRecord takes an int and adds it to the linked list.
func (sc *stringCollection) intToList(num int) {
	newNode := node{
		data: num,
		next: sc.head,
	}
	sc.head = &newNode
}

// // Reverse reverses the node position and utilizes printData to display the list.
func (sc *stringCollection) Reverse() {
	curr := sc.head
	var prev *node
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	sc.head = prev
}

func (sc stringCollection) printData() {
	toPrint := sc.head
	for toPrint != nil {
		fmt.Print(toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Println()
}

func main() {
	stringC := stringCollection{}
	stringC.intToList(1)
	stringC.intToList(4)
	stringC.intToList(9)
	stringC.Reverse()
	stringC.printData()
}
