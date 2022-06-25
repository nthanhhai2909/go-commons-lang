package sliceutils

import (
	"fmt"
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
		t.Errorf("IsEmpty func failed")
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
		t.Errorf("IsNotEmpty func failed")
	}
}

func TestMergeCollection(t *testing.T) {
	col1 := []int{1, 2, 3}
	col2 := []int{4, 5, 6}

	result := MergeSlices(col1, col2)
	if len(result) != 6 ||
		result[0] != 1 ||
		result[5] != 6 {
		t.Errorf("MergeSlices func failed")
	}
}

func TestMergeCollectionSortedCase1(t *testing.T) {
	col1 := []int{1, 3, 6}
	col2 := []int{2, 5, 4}

	result := MergeSlicesSorted(func(a, b int) bool { return a < b }, col1, col2)
	if len(result) != 6 ||
		result[0] != 1 ||
		result[5] != 6 {
		t.Errorf("MergeSlicesSorted func failed")
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
		t.Errorf("MergeSlicesSorted func failed in case sort asc by name")
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
		t.Errorf("MergeSlicesSorted func failed in case sort desc by name")
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
		t.Errorf("MergeSlicesSorted func failed in case sort desc by name")
	}
}

func TestSingletonSlice(t *testing.T) {
	slice := SingletonSlice(1)
	if len(slice) != 1 || slice[0] != 1 {
		t.Errorf("SingletonSlice func failed")
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3}
	Reverse(slice)
	if slice[0] != 3 ||
		slice[1] != 2 ||
		slice[2] != 1 {
		t.Errorf("Reverse func failed")
	}
}

func TestFirstCase1(t *testing.T) {
	item1 := First([]int{})
	if item1 != nil {
		t.Errorf("First func failed in case empty slice")
	}

	item2 := First([]int{1, 2, 3})
	if item2 == nil || item2.(int) != 1 {
		t.Errorf("First func failed to get first element")
	}
}

func TestFirstCase2(t *testing.T) {
	students := []student{{name: "a", age: 17}}
	firstStudent := First(students)
	if firstStudent == nil || firstStudent.(student).name != "a" || firstStudent.(student).age != 17 {
		t.Errorf("First func failed to get first element")
	}
}

func TestLastCase1(t *testing.T) {
	item1 := Last([]int{})
	if item1 != nil {
		t.Errorf("Last func failed in case empty slice")
	}

	item2 := Last([]int{1, 2, 3})
	if item2 == nil || item2.(int) != 3 {
		t.Errorf("Last func failed to get last element")
	}
}

func TestLastCase2(t *testing.T) {
	students := []student{{name: "a", age: 17}, {name: "b", age: 18}}
	lastStudent := Last(students)
	if lastStudent == nil || lastStudent.(student).name != "b" || lastStudent.(student).age != 18 {
		t.Errorf("Last func failed to get first element")
	}
}

func TestUniqueCase1(t *testing.T) {
	intSlice := []int{1, 2, 3, 2, 4}
	result := Unique(intSlice)

	if len(result) != 4 ||
		result[0] != 1 ||
		result[1] != 2 ||
		result[2] != 3 ||
		result[3] != 4 {
		t.Errorf("Unique func failed at test case 1")
	}
}

func TestUniqueCase2(t *testing.T) {
	strSlice := []string{"a", "b", "bb", "a", "b"}
	result := Unique(strSlice)

	if len(result) != 3 ||
		result[0] != "a" ||
		result[1] != "b" ||
		result[2] != "bb" {
		t.Errorf("Unique func failed at test case 1")
	}
}

func TestUniqueCase3(t *testing.T) {
	students := []student{{name: "a", age: 17}, {name: "b", age: 17}, {name: "a", age: 17}}
	result := Unique(students)
	fmt.Println(result)
	if len(result) != 2 ||
		result[0].name != "a" || result[0].age != 17 || result[1].name != "b" || result[1].age != 17 {
		t.Errorf("Unique func failed at test casse 2")
	}
}

func TestUniqueCase4(t *testing.T) {
	var students []student
	result := Unique(students)

	if len(result) != 0 {
		t.Errorf("Unique func failed at test casse 3")
	}
}

func TestIndexOf(t *testing.T) {
	strSlice := []string{"a", "b", "bb", "a", "b"}
	if IndexOf(strSlice, "bb") != 2 || IndexOf(strSlice, "ddd") != -1 {
		t.Errorf("IndexOf func failed at test case 1")
	}

	intSlice := []int{1, 2, 3, 4, 5}
	if IndexOf(intSlice, 4) != 3 || IndexOf(intSlice, 10) != -1 {
		t.Errorf("IndexOf func failed at test case 2")
	}

	students := []student{{name: "a", age: 17}, {name: "b", age: 18}, {name: "c", age: 19}}

	if IndexOf(students, student{name: "c", age: 19}) != 2 || IndexOf(students, student{name: "c", age: 30}) != -1 {
		t.Errorf("IndexOf func failed at test case 3")
	}
}

func TestContains(t *testing.T) {
	strSlice := []string{"a", "b", "bb", "a", "b"}
	if !Contains(strSlice, "bb") || Contains(strSlice, "ddd") {
		t.Errorf("Contains func failed at test case 1")
	}

	intSlice := []int{1, 2, 3, 4, 5}
	if !Contains(intSlice, 4) || Contains(intSlice, 10) {
		t.Errorf("Contains func failed at test case 2")
	}

	students := []student{{name: "a", age: 17}, {name: "b", age: 18}, {name: "c", age: 19}}

	if !Contains(students, student{name: "c", age: 19}) || Contains(students, student{name: "c", age: 30}) {
		t.Errorf("Contains func failed at test case 3")
	}
}

func TestNotContains(t *testing.T) {
	strSlice := []string{"a", "b", "bb", "a", "b"}
	if NotContains(strSlice, "bb") || !NotContains(strSlice, "ddd") {
		t.Errorf("NotContains func failed at test case 1")
	}

	intSlice := []int{1, 2, 3, 4, 5}
	if NotContains(intSlice, 4) || !NotContains(intSlice, 10) {
		t.Errorf("NotContains func failed at test case 2")
	}

	students := []student{{name: "a", age: 17}, {name: "b", age: 18}, {name: "c", age: 19}}

	if NotContains(students, student{name: "c", age: 19}) || !NotContains(students, student{name: "c", age: 30}) {
		t.Errorf("NotContains func failed at test case 3")
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
