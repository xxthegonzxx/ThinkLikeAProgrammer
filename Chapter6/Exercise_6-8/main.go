package main

import (
	"fmt"
	"sort"
)

type listNode struct {
	grade     int
	studentID int
	name      string
	next      *listNode
}

type studentCollection struct {
	head *listNode
}

func (sc *studentCollection) addRecord(n *listNode) {
	newNode := listNode{
		grade:     n.grade,
		studentID: n.studentID,
		name:      n.name,
		next:      sc.head,
	}
	sc.head = &newNode
}

func recurseRemoveRecord(n *listNode, ID int) {
	if n == nil {
		return
	}
	recurseRemoveRecord(n.next, ID)
	if n.studentID != ID {
		fmt.Printf("%d %d %s\n", n.grade, n.studentID, n.name)
	}
}

func (sc studentCollection) traverseList() {
	student := sc.head
	for student != nil {
		fmt.Printf("%d %d %s\n", student.grade, student.studentID, student.name)
		student = student.next
	}
}

func main() {
	studentList := studentCollection{}
	students := []listNode{
		{grade: 68, studentID: 1002, name: "Freddie"},
		{grade: 100, studentID: 1003, name: "Gonzalo"},
		{grade: 48, studentID: 1001, name: "Eric"},
		{grade: 78, studentID: 1004, name: "Sandra"},
		{grade: 98, studentID: 1005, name: "Irina"},
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i].grade < students[j].grade
	})
	for student := range students {
		studentList.addRecord(&students[student])
	}
	fmt.Println("By Grade:")
	studentList.traverseList()
	fmt.Println("-----")
	recurseRemoveRecord(studentList.head, 1001)
}
