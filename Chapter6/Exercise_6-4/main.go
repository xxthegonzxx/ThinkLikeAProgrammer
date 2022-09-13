package main

import "fmt"

func iterativeIsSorted(intArray []int, n int) bool {
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

func recursiveIsSorted(intArray []int, elementsInArray int) bool {
	if elementsInArray == 0 || elementsInArray == 1 {
		return true
	}
	lastElement := intArray[elementsInArray-1]
	if intArray[elementsInArray-2] <= lastElement {
		return recursiveIsSorted(intArray, elementsInArray-1)
	}
	return false
}

func main() {
	dupSorted := []int{1, 2, 3, 3, 4, 6}
	dupUnsorted := []int{5, 1, 1, 3, 4, 2, 2}
	fmt.Println(iterativeIsSorted(dupSorted, len(dupSorted)))
	fmt.Println(iterativeIsSorted(dupUnsorted, len(dupUnsorted)))
	fmt.Println(recursiveIsSorted(dupSorted, len(dupSorted)))
	fmt.Println(recursiveIsSorted(dupUnsorted, len(dupUnsorted)))
}
