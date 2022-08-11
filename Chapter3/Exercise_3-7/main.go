package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ARRAY_SIZE = 10
var intArray = make([]int, ARRAY_SIZE)
var modeCount = 1

func generateArr(n int) []int {
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		intArray[i] = rand.Int() % 10
	}
	return intArray
}

func findMode(arr []int) []int {
	mode := []int{}
	occurrences := map[int]int{}
	for _, v := range arr {
		occurrences[v]++
	}
	for number, count := range occurrences {
		fmt.Println(number, count)
		if count == modeCount {
			mode = append(mode, number)
		} else if count > modeCount {
			modeCount = count
			mode = []int{number}
		}
	}
	if modeCount == 1 {
		fmt.Println("No mode exists.")
	} else {
		fmt.Printf("The mode is %d and it appeared %d times.", mode, modeCount)
	}
	return mode
}

func main() {
	generateArr(ARRAY_SIZE)
	findMode(intArray)
}
