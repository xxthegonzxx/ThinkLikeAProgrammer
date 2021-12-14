package main

import "fmt"

func main() {
	for row := 0; row < 8; row = row + 2 {
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		for hashNum := 1; hashNum <= 8-row; hashNum++ {
			fmt.Print("#")
		}
		for i := 0; i < row/2; i++ {
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
