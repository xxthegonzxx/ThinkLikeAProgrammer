package main

import (
	"fmt"
	"sort"
)

const NUM_AGENTS = 3
const NUM_MONTHS = 12

var agentSumArray []int
var sales = [NUM_AGENTS][NUM_MONTHS]int{
	{1856, 498, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, 276, 32},
	{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, 7766, 2534},
	{23, 55, 67, 99, 265, 376, 232, 223, 4546, 564, 4544, 3434},
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

// findMedian, takes the number of months, and if even, finds the median of the
// two middle values. If odd, finds the median of the single middle value.
func findMedian(salesArray [NUM_AGENTS][NUM_MONTHS]int, NUM_MONTHS int) int {
	highestMedian := 0
	firstMiddle := (NUM_MONTHS / 2)
	secondMiddle := (NUM_MONTHS / 2)
	for month := range sales {
		if NUM_MONTHS%2 == 0 {
			firstMiddle := sales[month][firstMiddle-1]
			secondMiddle := sales[month][secondMiddle]
			agentSumArray = append(agentSumArray, firstMiddle, secondMiddle)
		} else {
			oddMedian := (NUM_MONTHS + 1) / 2
			oddMedian = sales[month][oddMedian-1]
			agentSumArray = append(agentSumArray, oddMedian)
		}
		agentMedian := arrayAverage(agentSumArray, len(agentSumArray))
		if agentMedian > highestMedian {
			highestMedian = agentMedian
		}
		agentSumArray = nil
	}
	return highestMedian
}

func main() {
	// First Sort the data
	for agent := 0; agent < NUM_AGENTS; agent++ {
		sort.Sort(orderedMonthlySales(sales[agent][:]))
	}
	// Then find the highest median sales
	fmt.Println("Highest Median:", findMedian(sales, NUM_MONTHS))
}
