package main

import "fmt"

func iterativeNumOcurrences(intArr []int, target int) int {
	count := 0
	for _, v := range intArr {
		if target == v {
			count++
		}
	}
	return count
}

func recursiveNumOccurrences(intArr []int, target int) int {
	size := len(intArr)
	if size == 0 {
		return 0
	}
	lastTarget := intArr[size-1]
	count := iterativeNumOcurrences(intArr[:size-1], target)
	if target == lastTarget {
		count++
	}
	return count
}

func main() {
	intArr := []int{5, 8, 7, 1, 8, 3, 7, 5, 7, 7, 3, 2, 5, 7, 0, 7}
	iterativeToPrint := iterativeNumOcurrences(intArr, 3)
	fmt.Println(iterativeToPrint)
	recursiveToPrint := recursiveNumOccurrences(intArr, 3)
	fmt.Println(recursiveToPrint)
}
