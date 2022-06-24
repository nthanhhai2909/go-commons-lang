package sliceutils

import (
	"testing"
)

type student struct {
	name string
	age  int
}

func TestIsEmpty(t *testing.T) {
	intCollection := make([]int, 0)
	strCollection := make([]string, 0)
	fCollection := make([]float32, 0)
	notEmptyCollection := []int{1, 2, 3}
	if !IsEmpty(intCollection) ||
		!IsEmpty(strCollection) ||
		!IsEmpty(fCollection) ||
		IsEmpty(notEmptyCollection) {
		t.Errorf("IsEmpty failed")
	}
}

func TestIsNotEmpty(t *testing.T) {
	intCollection := make([]int, 0)
	strCollection := make([]string, 0)
	fCollection := make([]float32, 0)
	notEmptyCollection := []int{1, 2, 3}
	if IsNotEmpty(intCollection) ||
		IsNotEmpty(strCollection) ||
		IsNotEmpty(fCollection) ||
		!IsNotEmpty(notEmptyCollection) {
		t.Errorf("IsNotEmpty failed")
	}
}

func TestMergeCollection(t *testing.T) {
	col1 := []int{1, 2, 3}
	col2 := []int{4, 5, 6}

	result := MergeSlices(col1, col2)
	if len(result) != 6 ||
		result[0] != 1 ||
		result[5] != 6 {
		t.Errorf("MergeSlices failed")
	}
}

func TestMergeCollectionSortedCase1(t *testing.T) {
	col1 := []int{1, 3, 6}
	col2 := []int{2, 5, 4}

	result := MergeSlicesSorted(func(a, b int) bool { return a < b }, col1, col2)
	if len(result) != 6 ||
		result[0] != 1 ||
		result[5] != 6 {
		t.Errorf("MergeSlicesSorted failed")
	}
}

func TestMergeCollectionSortedCase2(t *testing.T) {
	student1, student2 := studentData()
	sortedAscByName := MergeSlicesSorted(func(a, b student) bool {
		return a.name < b.name
	}, student1, student2)

	if len(sortedAscByName) != 6 ||
		sortedAscByName[0].name != "a" && sortedAscByName[0].age != 24 ||
		sortedAscByName[1].name != "a" && sortedAscByName[1].age != 34 ||
		sortedAscByName[5].name != "f" && sortedAscByName[1].age != 21 {
		t.Errorf("MergeSlicesSorted failed in case sort asc by name")
	}
}

func TestMergeCollectionSortedCase3(t *testing.T) {
	student1, student2 := studentData()
	sortedAscByAge := MergeSlicesSorted(func(a, b student) bool {
		return a.age < b.age
	}, student1, student2)

	if len(sortedAscByAge) != 6 ||
		sortedAscByAge[0].name != "c" && sortedAscByAge[0].age != 19 ||
		sortedAscByAge[1].name != "f" && sortedAscByAge[1].age != 24 ||
		sortedAscByAge[5].name != "b" && sortedAscByAge[5].age != 54 {
		t.Errorf("MergeSlicesSorted failed in case sort desc by name")
	}
}

func TestMergeCollectionSortedCase4(t *testing.T) {
	student1, student2 := studentData()
	sortedDescByAge := MergeSlicesSorted(func(a, b student) bool {
		return a.age > b.age
	}, student1, student2)

	if len(sortedDescByAge) != 6 ||
		sortedDescByAge[5].name != "c" && sortedDescByAge[5].age != 19 ||
		sortedDescByAge[0].name != "b" && sortedDescByAge[0].age != 54 {
		t.Errorf("MergeSlicesSorted failed in case sort desc by name")
	}
}

func studentData() (student1, student2 []student) {
	student1 = make([]student, 3)
	student2 = make([]student, 3)
	student1[0] = student{name: "a", age: 24}
	student1[1] = student{name: "f", age: 21}
	student1[2] = student{name: "d", age: 45}

	student2[0] = student{name: "a", age: 34}
	student2[1] = student{name: "c", age: 19}
	student2[2] = student{name: "b", age: 54}
	return student1, student2
}
