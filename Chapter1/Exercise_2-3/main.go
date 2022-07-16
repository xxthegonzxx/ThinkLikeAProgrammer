package main

import "fmt"

func main() {
	row := 0
	totalCharacterNum := 14
	var hashCount = row + 1
	var sideSpaceCount = row
	var spaceCount = totalCharacterNum - (hashCount * 2) - (sideSpaceCount * 2)
	fmt.Print("#")
	for space := 0; space < spaceCount; space++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Print("\n")
	row++
	for spaceCount > 0 {
		hashCount = row + 1
		sideSpaceCount = row
		spaceCount = totalCharacterNum - (hashCount * 2) - (sideSpaceCount * 2)
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
		row++
	}
	for row := 3; row >= 0; row-- {
		hashCount = row + 1
		sideSpaceCount = row
		spaceCount = totalCharacterNum - (hashCount * 2) - (sideSpaceCount * 2)
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
