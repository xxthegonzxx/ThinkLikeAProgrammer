package main

import "fmt"

func main() {
	for row := 0; row <= 3; row++ {
		totalCharacterCount := 8
		spaceCount := row + row
		hashCount := totalCharacterCount - spaceCount
		for space := 0; space < spaceCount/2; space++ {
			fmt.Print(" ")
		}
		for hash := 0; hash < hashCount; hash++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
