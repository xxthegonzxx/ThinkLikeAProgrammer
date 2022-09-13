package main

import "fmt"

func iterativeOddParity(strArray []string) bool {
	count := 0
	for _, v := range strArray {
		if v == "1" {
			count++
		}
	}
	if count%2 == 0 {
		return false
	} else {
		return true
	}
}

func recursiveOddParity(strArray []string, size int) bool {
	if size == 0 {
		return false
	}
	prevResult := recursiveOddParity(strArray, size-1)
	if prevResult {
		return strArray[size-1] != "1"
	}
	return strArray[size-1] == "1"
}

func main() {
	strArray := []string{"1", "0", "1", "1", "0"}
	printiterativeOddParity := iterativeOddParity(strArray)
	fmt.Println(printiterativeOddParity)
	printRecursiveOddParity := recursiveOddParity(strArray, len(strArray))
	fmt.Println(printRecursiveOddParity)
}
