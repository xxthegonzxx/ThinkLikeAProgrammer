package main

import "fmt"

var intArray = []int{15, 11, 66, 39, 2395, 11, 11, 11, 33, 34, 33, 76, 67, 33, 11, 34, 90, 3, 1, 2, 8, 4, 2, 3, 6, 7, 0, 2}

func main() {
	highestFrequency := 0
	mostFrequent := 0
	occurrences := map[int]int{}
	for _, v := range intArray {
		occurrences[v]++
	}

	for number, count := range occurrences {
		if count >= highestFrequency {
			highestFrequency = count
			mostFrequent = number
		}
	}
	fmt.Printf("%d appears %d times\n", mostFrequent, highestFrequency)
}
