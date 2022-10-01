package main

import "fmt"

type studentRecord struct {
	name         string
	id           int
	grade        int
	extraOptions map[string]any
}

func (sr *studentRecord) getTitle() any {
	option := sr.extraOptions["Title"]
	if option == nil {
		return nil
	}
	return option
}

func (sr *studentRecord) getYear() any {
	option := sr.extraOptions["Year"]
	if option == nil {
		return nil
	}
	return option
}

func (sr *studentRecord) getAudit() any {
	option := sr.extraOptions["Audit"]
	if option == nil {
		return nil
	}
	return option
}

func (sr *studentRecord) displayRecord() studentRecord {
	newRecord := studentRecord{
		name:         sr.name,
		id:           sr.id,
		grade:        sr.grade,
		extraOptions: sr.extraOptions,
	}
	return newRecord
}

func main() {
	studentOptions1 := make(map[string]any)
	studentOptions1["Title"] = "My Title."
	studentOptions1["Year"] = 2020
	studentOptions1["Audit"] = true
	Student1 := studentRecord{
		name:         "Gonzo",
		id:           1001,
		grade:        100,
		extraOptions: studentOptions1,
	}
	studentOptions2 := make(map[string]any)
	studentOptions2["Year"] = 2020
	student2 := studentRecord{
		name:         "Ted",
		id:           1002,
		grade:        89,
		extraOptions: studentOptions2,
	}

	fmt.Println(Student1.displayRecord())
	fmt.Println(student2.displayRecord())
}
