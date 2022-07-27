package main

import (
	"fmt"
	"sort"
)

type student struct {
	grade     int
	studentID int
	name      string
}

// byGrade implements sort.Interface for []student based on
// the grade field.
type byGrade []student

func (g byGrade) Len() int { return len(g) }

func (g byGrade) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

func (g byGrade) Less(i, j int) bool { return g[i].grade < g[j].grade }

// byStudentId implements sort.Interface for []student based on
// the studentID field.
type byStudentId []student

func (g byStudentId) Len() int { return len(g) }

func (g byStudentId) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

func (g byStudentId) Less(i, j int) bool { return g[i].studentID < g[j].studentID }

func main() {
	studentArray := []student{
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

	sort.Slice(studentArray, func(i, j int) bool {
		return studentArray[i].grade < studentArray[j].grade
	})
	fmt.Println("Sorting by Slice method and comparator function:\nBy Grade: ", studentArray)

	sort.Slice(studentArray, func(i, j int) bool {
		return studentArray[i].studentID < studentArray[j].studentID
	})
	fmt.Println("By StudentID: ", studentArray)

	sort.Sort(byGrade(studentArray))
	fmt.Println("--------------------")
	fmt.Println("Sorting by the Sort method using sort.Interface\nBy Grade", studentArray)

	sort.Sort(byStudentId(studentArray))
	fmt.Println("By StudentID:", studentArray)

}
