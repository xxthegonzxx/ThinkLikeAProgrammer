package main

import (
	"fmt"
	"sort"
)

type studentRecord struct {
	name  string
	id    int
	grade int
}

type studentCollection []studentRecord

func interpolationSearch(studentArr studentCollection, studentID int) studentRecord {
	mid := 0
	low := 0
	high := len(studentArr) - 1
	for studentArr[low].id < studentID && studentArr[high].id > studentID {
		mid = low + ((studentID-studentArr[low].id)*(high-low))/(studentArr[high].id-studentArr[low].id)
		if studentArr[mid].id < studentID {
			low = mid + 1
		} else if studentArr[mid].id > studentID {
			high = mid - 1
		} else {
			return studentArr[mid]
		}
	}
	if studentArr[low].id == studentID {
		return studentArr[low]
	} else if studentArr[high].id == studentID {
		return studentArr[high]
	}
	return studentRecord{"RecordNotFound", 0, 0}
}

func main() {
	studentArray := studentCollection{
		studentRecord{"Gonzo", 1004, 100},
		studentRecord{"Eric", 1005, 70},
		studentRecord{"Irina", 1003, 80},
		studentRecord{"Ravi", 1001, 78},
		studentRecord{"Patrick", 1002, 88},
		studentRecord{"Gandalf", 1006, 100},
		studentRecord{"Aragorn", 1007, 73},
		studentRecord{"Galadriel", 1008, 99},
		studentRecord{"Frodo", 1009, 76},
		studentRecord{"Samwise", 1010, 74},
		studentRecord{"Sauron", 1011, 89},
		studentRecord{"Legolas", 1012, 95},
		studentRecord{"Arwen", 1013, 98},
		studentRecord{"Isildur", 1014, 88},
		studentRecord{"Golum", 1015, 67},
	}
	sort.Slice(studentArray, func(i, j int) bool {
		return studentArray[i].id < studentArray[j].id
	})
	findRecord := interpolationSearch(studentArray, 1011)
	fmt.Println(findRecord)
}
