package main

import "fmt"

type node struct {
	data rune
	next *node
}

type stringCollection struct {
	head   *node
	length int
}

// addRecord takes a string and adds it to the linked list.
func (sc *stringCollection) addRecord(str string) {
	for _, v := range str {
		newNode := node{
			data: rune(v),
			next: sc.head,
		}
		sc.head = &newNode
	}
}

// removeChars removes a section of characters from a string based on the position and length.
func (sc *stringCollection) removeChars(pos, charLength int) {
	startPos := 0
	current := sc.head
	var prev *node

	for startPos != pos && current != nil && current.next != nil {
		prev = current
		current = current.next
		startPos++
	}

	for startPos < pos+charLength {
		if current != nil {
			current = current.next
		}
		startPos++
	}
	prev.next = current
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
}

func (sc stringCollection) printData() {
	toPrint := sc.head
	for toPrint != nil {
		metaData := string(toPrint.data)
		fmt.Print(metaData)
		toPrint = toPrint.next
	}
	fmt.Println()
}

func main() {
	stringC := stringCollection{}
	stringC.addRecord("Golang")
	stringC.Reverse()
	stringC.removeChars(2, 4)
	stringC.printData()
}
