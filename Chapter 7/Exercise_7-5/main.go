package main

import (
	"container/list"
	"fmt"
)

type studentRecord struct {
	name  string
	id    int
	grade int
}

type studentCollection struct {
	students *list.List
}

func (sc *studentCollection) addRecord(student *studentRecord) {
	newNode := studentRecord{
		grade: student.grade,
		id:    student.id,
		name:  student.name,
	}
	sc.students.PushFront(&newNode)
}

func (sc studentCollection) searchById(id int) bool {
	elem := sc.students.Front()

	for elem != nil {
		student := elem.Value.(*studentRecord)
		if student.id == id {
			fmt.Println("Name:", student.name, "\tGrade:", student.grade, "\tStudent ID:", student.id)
			return true
		}
		elem = elem.Next()
	}
	fmt.Println("Student not found.")
	return false
}

func main() {
	sc := studentCollection{
		students: list.New(),
	}
	studentList := []studentRecord{
		{name: "Gonzo", id: 1004, grade: 100},
		{name: "Eric", id: 1005, grade: 70},
		{name: "Irina", id: 1003, grade: 80},
		{name: "Ravi", id: 1001, grade: 78},
		{name: "Patrick", id: 1002, grade: 88},
		{name: "Gandalf", id: 1006, grade: 100},
		{name: "Aragorn", id: 1007, grade: 73},
		{name: "Galadriel", id: 1008, grade: 99},
		{name: "Frodo", id: 1009, grade: 76},
		{name: "Samwise", id: 1010, grade: 74},
	}

	for record := range studentList {
		sc.addRecord(&studentList[record])
	}
	sc.searchById(1008)

}
