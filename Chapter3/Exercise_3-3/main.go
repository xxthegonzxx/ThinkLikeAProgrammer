package main

import (
	"fmt"
)

func recursiveIsSorted(intArray []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if intArray[i] < intArray[i-1] {
			return false
		}
	}
	return true
}

func main() {
	dupSorted := []int{20, 21, 45, 89, 89, 90}
	dupUnsorted := []int{20, 20, 78, 98, 99, 97}
	fmt.Println(recursiveIsSorted(dupSorted, len(dupSorted)))
	fmt.Println(recursiveIsSorted(dupUnsorted, len(dupUnsorted)))
}
