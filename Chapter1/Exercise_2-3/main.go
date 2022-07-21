package main

import "fmt"

func main() {
	totalCharacterNum := 14
	for row := 0; row <= 3; row++ {
		hashCount := row + 1
		sideSpaceCount := row
		spaceCount := totalCharacterNum - (hashCount * 2) - (sideSpaceCount * 2)
		for sideSpace := 0; sideSpace < sideSpaceCount; sideSpace++ {
			fmt.Print(" ")
		}
		for hashMark := 0; hashMark < hashCount; hashMark++ {
			fmt.Print("#")
		}
		for space := 0; space < spaceCount; space++ {
			fmt.Print(" ")
		}
		for hashMark := 0; hashMark < hashCount; hashMark++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
	for row := 3; row >= 0; row-- {
		hashCount := row + 1
		sideSpaceCount := row
		spaceCount := totalCharacterNum - (hashCount * 2) - (sideSpaceCount * 2)
		for sideSpace := 0; sideSpace < sideSpaceCount; sideSpace++ {
			fmt.Print(" ")
		}
		for hashMark := 0; hashMark < hashCount; hashMark++ {
			fmt.Print("#")
		}
		for space := 0; space < spaceCount; space++ {
			fmt.Print(" ")
		}
		for hashMark := 0; hashMark < hashCount; hashMark++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
