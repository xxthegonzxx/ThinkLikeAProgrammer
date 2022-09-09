package main

import "fmt"

func iterativeSumNumbers(numArray []int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += numArray[i]
	}
	return sum
}

func recursiveSumNumbers(numArray []int, size int) int {
	if size == 0 {
		return 0
	}
	lastNumber := numArray[size-1]
	allButLastSum := recursiveSumNumbers(numArray, size-1)
	return lastNumber + allButLastSum
}

func main() {
	arr := []int{5, 5, 10, 2, 4, 8, 6}
	printIterativeSum := iterativeSumNumbers(arr, len(arr))
	fmt.Println(printIterativeSum)
	printRecursiveSum := recursiveSumNumbers(arr, len(arr))
	fmt.Println(printRecursiveSum)
}
