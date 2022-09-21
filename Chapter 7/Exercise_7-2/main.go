package main

import (
	"container/list"
	"fmt"
)

type student struct {
	grade     int
	studentID int
	name      string
}

type studentCollection struct {
	students *list.List
}

func (sc studentCollection) averageRecord() int {
	elem := sc.students.Front()

	if elem == nil {
		return 0
	}

	sum := 0
	studentCount := 0

	for elem != nil {
		sum += elem.Value.(*student).grade
		studentCount++
		elem = elem.Next()
	}

	return sum / studentCount
}

func (sc *studentCollection) addRecord(firstName string, studentGrade, id int) {
	newNode := student{
		grade:     studentGrade,
		studentID: id,
		name:      firstName,
	}
	sc.students.PushFront(&newNode)
}

func (sc studentCollection) printRecords() {
	elem := sc.students.Front()

	fmt.Println("Printing student list.")
	for elem != nil {
		student := elem.Value.(*student)
		fmt.Println("Name:", student.name, "\tGrade:", student.grade, "\tStudent ID:", student.studentID)
		elem = elem.Next()
	}
}

func main() {
	sc := studentCollection{
		students: list.New(),
	}
	sc.addRecord("Gonzo", 100, 1001)
	sc.addRecord("Eric", 80, 1002)
	sc.addRecord("Gabriela", 78, 1003)
	sc.printRecords()
	printSum := sc.averageRecord()
	fmt.Println("The average is:", printSum)
}
