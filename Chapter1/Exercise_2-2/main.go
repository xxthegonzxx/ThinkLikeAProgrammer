package main

import "fmt"

func main() {
	for row := 6; row >= 0; row -= 2 {
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		for hashNum := 0; hashNum <= 7-row; hashNum++ {
			fmt.Print("#")
		}
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	for row := 0; row < 8; row += 2 {
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		for hashNum := 0; hashNum <= 7-row; hashNum++ {
			fmt.Print("#")
		}
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
