package main

import (
	"fmt"
	"sort"
)

type firstStudentPolicy int

const (
	higherGrade firstStudentPolicy = iota
	lowerStudentNumber
	nameComestFirst
)

type studentRecord struct {
	grade     int
	studentID int
	name      string
}

type listNode struct {
	studentData studentRecord
	next        *listNode
}

type studentCollection struct {
	head               *listNode
	firstStudentPolicy func(r1, r2 studentRecord) bool
}

func (sc *studentCollection) addRecord(student studentRecord) {
	newNode := listNode{
		studentData: student,
		next:        sc.head,
	}
	sc.head = &newNode
}

func (sc *studentCollection) deleteRecord(name string) {
	student := sc.head
	for student != nil {
		if student.studentData.name != name {
			fmt.Printf("%d %d %s\n", student.studentData.grade, student.studentData.studentID, student.studentData.name)
		}
		student = student.next
	}
}

func (sc studentCollection) traverseList() {
	student := sc.head
	for student != nil {
		fmt.Printf("%d %d %s\n", student.studentData.grade, student.studentData.studentID, student.studentData.name)
		student = student.next
	}
}

var firstStudentPolicies = map[firstStudentPolicy]func(r1, r2 studentRecord) bool{
	higherGrade: func(r1, r2 studentRecord) bool {
		return r2.grade > r1.grade
	},
	lowerStudentNumber: func(r1, r2 studentRecord) bool {
		return r2.studentID < r1.studentID
	},
	nameComestFirst: func(r1, r2 studentRecord) bool {
		return r2.name < r1.name
	},
}

func (sc *studentCollection) setFirstStudentPolicy(policy firstStudentPolicy) {
	sc.firstStudentPolicy = firstStudentPolicies[policy]
}

func (sc *studentCollection) getFirstStudent() *listNode {
	if sc.firstStudentPolicy == nil {
		return nil
	}

	node := sc.head
	firstStudent := node

	for node != nil && node.next != nil {
		if sc.firstStudentPolicy(node.studentData, node.next.studentData) {
			firstStudent = node.next
		}
		node = node.next
	}

	return firstStudent
}

func (sc *studentCollection) byStudentId(id int) *listNode {
	node := sc.head

	for node.studentData.studentID != id {
		node = node.next

		if node == nil {
			return nil
		}
	}

	return node
}

func main() {
	studentList := studentCollection{}
	students := []studentRecord{
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
		studentList.addRecord(students[student])
	}

	for _, policy := range []firstStudentPolicy{0, 1, 2} {
		studentList.setFirstStudentPolicy(policy)
		s := studentList.getFirstStudent()
		fmt.Printf(
			"First student by Policy %d: %s (ID: %d, grade: %d)\n",
			policy,
			s.studentData.name,
			s.studentData.studentID,
			s.studentData.grade,
		)
	}
}
