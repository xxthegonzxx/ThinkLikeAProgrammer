package main

import "fmt"

func main() {
	totalCharacterCount := 8
	for row := 3; row >= 0; row-- {
		drawDiamond(row, totalCharacterCount)
	}
	for row := 0; row <= 3; row++ {
		drawDiamond(row, totalCharacterCount)
	}
}

func drawDiamond(row int, totalCharacterCount int) {
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
