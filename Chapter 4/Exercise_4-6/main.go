package main

import (
	"fmt"
)

type node struct {
	data rune
	next *node
}

type stringCollection struct {
	head *node
}

// append takes a string and a character and appends the character to the end of the string.
func (sc *stringCollection) append(str string, char rune) {
	for _, v := range str {
		newNode := node{
			data: rune(v),
			next: sc.head,
		}
		sc.head = &newNode
	}
	singleCharNode := node{
		data: char,
		next: sc.head,
	}
	sc.head = &singleCharNode

}

// characterAt takes a number as a parameter and returns the character at that
// position in the string (with the first character(node) numbered zero).
func (sc stringCollection) characterAt(position int) string {
	location := sc.head
	locationPos := 0
	for position != locationPos {
		location = location.next
		locationPos++
	}
	return string(location.data)
}

// printData traverses the linked list and prints out each node in the list.
func (sc stringCollection) printData() {
	traverseHead := sc.head
	for traverseHead != nil {
		fmt.Print(string(traverseHead.data))
		traverseHead = traverseHead.next
	}
	fmt.Println()
}

// reverse reverses the node position and utilizes printData to display the list.
func (sc *stringCollection) reverse() {
	curr := sc.head
	var prev *node

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	sc.head = prev
	sc.printData()
}

func main() {
	stringC := stringCollection{}
	stringC.append("golang", '!')
	stringC.reverse()
	fmt.Println(stringC.characterAt(5))
}
