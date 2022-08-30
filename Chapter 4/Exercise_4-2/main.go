package main

import (
	"fmt"
	"log"
)

// func subString takes three parameters: an arrayString, a starting position integer,
// and an integer length of characters. The function returns a new dynamically
// allocated string slice. This string slice contains the characters in the original string,
// starting at the specified position for the specified length.
func subString(arrayString []string, startPos int, length int) []string {
	newArrString := []string{}
	if length > len(arrayString) {
		log.Fatal("Length out of bounds.")
	}
	newArrString = append(newArrString, arrayString[startPos-1:startPos-1+length]...)
	return newArrString
}

func main() {
	arrayString := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(subString(arrayString, 3, 4))
}
