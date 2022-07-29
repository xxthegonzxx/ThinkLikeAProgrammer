// To find the median:
// Arrange the data points from smallest to largest.
// If the number of data points is odd, the median is the middle data point in the list.
// If the number of data points is even, the median is the average of the two middle data points in the list.
package main

import (
	"fmt"
	"sort"
)

const NUM_AGENTS = 3
const NUM_MONTHS = 12

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

func main() {
	var agentSumArray []int
	highestMedian := 0
	// First Sort the data
	for agent := 0; agent < NUM_AGENTS; agent++ {
		sort.Sort(orderedMonthlySales(sales[agent][:]))
		for month := 0; month < NUM_MONTHS; month++ {
			if NUM_MONTHS%2 == 0 {
				firstMiddle := (NUM_MONTHS / 2)
				firstMiddle = sales[agent][firstMiddle-1]
				secondMiddle := (NUM_MONTHS / 2)
				secondMiddle = sales[agent][secondMiddle]
				agentSumArray = append(agentSumArray, firstMiddle, secondMiddle)
			} else {
				oddMedian := (NUM_MONTHS + 1) / 2
				oddMedian = sales[agent][oddMedian-1]
				agentSumArray = append(agentSumArray, oddMedian)
			}
		}
		agentMedian := arrayAverage(agentSumArray, len(agentSumArray))
		if agentMedian > highestMedian {
			highestMedian = agentMedian
		}
		agentSumArray = nil
	}
	fmt.Println("Highest Median:", highestMedian)
}
