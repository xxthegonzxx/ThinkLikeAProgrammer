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

// Append takes a string and a character and appends the character to the end of the string.
func (sc *stringCollection) Append(str string, char rune) {
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

// Concatenate takes two strings and appends the characters of the second string to the first.
func (sc *stringCollection) Concatenate(s1, s2 string) string {
	for _, v := range s1 {
		s1Node := node{
			data: rune(v),
			next: sc.head,
		}
		sc.head = &s1Node
	}
	for _, v := range s2 {
		s2Node := node{
			data: rune(v),
			next: sc.head,
		}
		sc.head = &s2Node
	}
	return string(sc.head.data)
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

// Reverse reverses the node position and utilizes printData to display the list.
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
	sc.printData()
}

func main() {
	stringC := stringCollection{}
	stringC.Concatenate("Golang", "Rocks!")
	stringC.Reverse()
}
