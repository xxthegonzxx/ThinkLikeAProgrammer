package main

import (
	"fmt"
	"sort"
)

var studentArray = []student{
	{87, 10001, "Fred"},
	{28, 10002, "Tom"},
	{100, 10003, "Gonzalo"},
	{78, 10004, "Sasha"},
	{84, 10005, "Erin"},
	{98, 10006, "Humberto"},
	{75, 10007, "Melinda"},
	{70, 10008, "Holly"},
	{81, 10009, "Deshawn"},
	{68, 10010, "Aretha"},
}
var NUM_STUDENTS = len(studentArray)
var medianArray []int

type student struct {
	grade     int
	studentID int
	name      string
}

// arrayAverage finds the average by taking the sum and dividing by the size of the array.
func arrayAverage(intArray []int, ARRAY_SIZE int) int {
	sum := 0
	for i := 0; i < ARRAY_SIZE; i++ {
		sum += intArray[i]
	}
	average := sum / ARRAY_SIZE
	return average
}

// findQuartile, finds the quartiles of the student slice passed into it.
// It uses several formulas to find the 1st, 2nd, and 3rd quartiles and returns
// 3 integers representing each quartile.
func findQuartile(students []student, NUM_STUDENTS int) (int, int, int) {
	firstQuartileFormula := (NUM_STUDENTS + 1) / 4
	thirdQuartileFormula := (firstQuartileFormula * 3)
	firstMiddleFormula := (NUM_STUDENTS / 2)
	secondMiddleFormula := (NUM_STUDENTS / 2)
	oddMedianFormula := (NUM_STUDENTS + 1) / 2
	if NUM_STUDENTS%2 == 0 {
		firstMiddle := students[firstMiddleFormula-1]
		secondMiddle := students[secondMiddleFormula]
		medianArray = append(medianArray, firstMiddle.grade, secondMiddle.grade)
	} else {
		oddMedian := students[oddMedianFormula-1]
		medianArray = append(medianArray, oddMedian.grade)
	}
	firstQuartile := students[firstQuartileFormula-1]
	secondQuartile := arrayAverage(medianArray, len(medianArray))
	thirdQuartile := students[thirdQuartileFormula]

	return firstQuartile.grade, secondQuartile, thirdQuartile.grade
}

func main() {
	sort.Slice(studentArray, func(i, j int) bool {
		return studentArray[i].grade < studentArray[j].grade
	})
	fmt.Println(studentArray)
	first, second, third := findQuartile(studentArray, NUM_STUDENTS)
	fmt.Printf("Receive a score of %d to do as well or better than 25%% your peers.\n", first)
	fmt.Printf("Receive a score of %d to do as well or better than 50%% your peers.\n", second)
	fmt.Printf("Receive a score of %d to do as well or better than 75%% your peers.\n", third)

}
