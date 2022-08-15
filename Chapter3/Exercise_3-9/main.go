package main

import (
	"fmt"
	"sort"
)

const NUM_AGENTS = 3
const NUM_MONTHS = 12

var sales = [NUM_AGENTS][NUM_MONTHS]int{
	{-1, -1, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, 276, 32},
	{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, 7766, -1},
	{23, 55, 67, 99, 265, 376, 232, 223, -1, -1, -1, -1},
}

// orderedMonthlySales sorts the data from smallest to largest
type orderedMonthlySales []int

func (a orderedMonthlySales) Len() int           { return len(a) }
func (a orderedMonthlySales) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a orderedMonthlySales) Less(i, j int) bool { return a[i] < a[j] }

// arrayAverage finds the average by taking the sum and dividing by the size of the array.
func arrayAverage(intArray []int, ARRAY_SIZE int) int {
	sum := 0
	for i := 0; i < ARRAY_SIZE; i++ {
		sum += intArray[i]
	}
	average := sum / ARRAY_SIZE
	return average
}

// findHighestMedian, takes the number of months and, depending on whether the months
// are odd or even, finds the median values and returns the highest median value.
func findHighestMedian(salesArray [NUM_AGENTS][NUM_MONTHS]int, NUM_MONTHS int) int {
	highestMedian := 0
	activeMonths := []int{}
	medianArray := []int{}
	for _, month := range sales {
		for sale := range month {
			if month[sale] == -1 {
				continue
			} else {

			}
			activeMonths = append(activeMonths, month[sale])
		}
		activeNumMonths := len(activeMonths)
		findActiveMiddleValue := (activeNumMonths / 2)
		findActiveOddMedian := (activeNumMonths + 1) / 2
		if activeNumMonths%2 == 0 {
			firstMedianVal := activeMonths[findActiveMiddleValue-1]
			secondMedianVal := activeMonths[findActiveMiddleValue]
			medianArray = append(medianArray, firstMedianVal, secondMedianVal)
		} else {
			oddActiveMedian := activeMonths[findActiveOddMedian-1]
			medianArray = append(medianArray, oddActiveMedian)
		}
		agentMedian := arrayAverage(medianArray, len(medianArray))
		fmt.Println("Agent Median Sales:", agentMedian)
		if agentMedian > highestMedian {
			highestMedian = agentMedian
		}
		activeMonths = nil
	}
	return highestMedian
}

func main() {
	// First Sort the data
	for agent := 0; agent < NUM_AGENTS; agent++ {
		sort.Sort(orderedMonthlySales(sales[agent][:]))
	}
	// Then find the highest median sales
	fmt.Println("Highest Median:", findHighestMedian(sales, NUM_MONTHS))
}
