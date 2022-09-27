package main

import (
	"fmt"
	"time"
)

const arraySize = 100

// Hash Table Structure
type hashTable struct {
	array [arraySize]*bucket
}

// BucketNode Struct
type bucketNode struct {
	name  string
	id    int
	grade int
	next  *bucketNode
}

// Bucket Structure (Linked List)
type bucket struct {
	head *bucketNode
}

// Insert will accept a student as parameter and add student to Hash Table.
func (h *hashTable) Insert(student bucketNode) {
	index := hash(student.id)
	h.array[index].insert(student.name, student.id, student.grade)
}

// Search will accept a student as a parameter and will return true if found in Hash Table.
func (h *hashTable) Search(studentId int) bool {
	index := hash(studentId)
	return h.array[index].search(studentId)
}

// insert
func (student *bucket) insert(studentName string, studentId, studentGrade int) {
	newNode := bucketNode{
		name:  studentName,
		id:    studentId,
		grade: studentGrade,
		next:  student.head,
	}
	student.head = &newNode
}

// search
func (bucket *bucket) search(studentId int) bool {
	current := bucket.head
	for current != nil {
		if studentId == current.id {
			fmt.Println("Student:", current.name, "\ngrade:", current.grade, "\nId:", current.id)
			return true
		}
		current = current.next
	}
	return false
}

// hash
func hash(id int) int {
	currentTime := time.Now()
	sum := 0
	sum += currentTime.Nanosecond() * id
	return sum % arraySize
}

// initFunc
func Init() *hashTable {
	result := &hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	studentList := []bucketNode{
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
	for student := range studentList {
		hashTable.Insert(studentList[student])
	}
	fmt.Println("Student Found?", hashTable.Search(1008))

}
