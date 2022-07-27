package main

import "fmt"

func main() {
	totalCharCount := 14
	for alphaHash := 0; alphaHash <= totalCharCount-1; alphaHash++ {
		fmt.Print("#")
	}
	fmt.Print("\n")
	for row := 0; row <= 6; row++ {
		drawForI(row, totalCharCount)
	}
	for row := 6; row >= 0; row-- {
		drawForI(row, totalCharCount)
	}
	for omegaHash := 0; omegaHash <= totalCharCount-1; omegaHash++ {
		fmt.Print("#")
	}
	fmt.Print("\n")
}

func drawForI(row int, totalCharCount int) {
	sideSpaceCount := row * 2
	hashCount := 2
	spaceCount := totalCharCount - sideSpaceCount - hashCount
	for sideSpace := 0; sideSpace < sideSpaceCount/2; sideSpace++ {
		fmt.Print(" ")
	}
	for hash := 0; hash < hashCount/2; hash++ {
		fmt.Print("#")
	}
	for space := 0; space < spaceCount; space++ {
		fmt.Print(" ")
	}
	for hash := 0; hash < hashCount/2; hash++ {
		fmt.Print("#")
	}
	fmt.Print("\n")
}
