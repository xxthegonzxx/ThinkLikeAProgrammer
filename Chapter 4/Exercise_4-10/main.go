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

func (sc *stringCollection) sumList(list1, list2 stringCollection) int {
	sum := 0
	headOne := list1.head
	headTwo := list2.head
	for headOne != nil && headTwo != nil {
		val1 := headOne.data
		headOne = headOne.next
		val2 := headTwo.data
		headTwo = headTwo.next
		sum += val1 + val2
	}
	return sum
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
	linkedListOne := stringCollection{}
	linkedListOne.intToList(5)
	linkedListOne.intToList(2)
	linkedListOne.intToList(2)
	linkedListOne.Reverse()
	linkedListTwo := stringCollection{}
	linkedListTwo.intToList(5)
	linkedListTwo.intToList(2)
	linkedListTwo.intToList(2)
	linkedListSum := stringCollection{}
	printSum := linkedListSum.sumList(linkedListOne, linkedListTwo)
	fmt.Println(printSum)
}
