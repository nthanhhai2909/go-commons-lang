package sliceutils

import (
	"go-commons-lang/types"
	"sort"
)

const (
	defaultCapacity = 10
	indexNotFound   = -1
)

/**
 * Check if a slice is empty or not.
 * @examples:
 * IsEmpty(make([]int, 0)) = true
 * IsEmpty([]int{1, 2, 3})) = false
 *
 * @parameters the slice to check
 * @returns {@code true} if the string is empty
 */

func IsEmpty[T any](slice []T) bool {
	if len(slice) == 0 {
		return true
	}
	return false
}

/**
 * Check if a slice is not empty.
 * @examples:
 * IsNotEmpty(make([]int, 0)) = false
 * IsNotEmpty([]int{1, 2, 3})) = true
 *
 * @parameters the slice to check
 * @returns {@code true} if the string is not empty
 */

func IsNotEmpty[T any](slice []T) bool {
	return !IsEmpty(slice)
}

/**
 * Merge all slice into a slice
 * @examples:
 * MergeSlices([]int{1, 2, 3}, []int{4, 5, 6}) = []int{1, 2, 3, 4, 5, 6}
 *
 * @parameters a list of slice
 * @returns Slice which have all element of all slice
 */

func MergeSlices[T any](slices ...[]T) []T {
	result := make([]T, 0, defaultCapacity)
	for _, coll := range slices {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	return result
}

/**
 * Merge all slice into a slice and then sort
 * @examples:
 * MergeSlicesSorted([]int{1, 3, 6}, []int{2, 5, 4}}) = []int{1, 2, 3, 4, 5, 6}
 *
 * @parameters a list of slice
 * @returns A sorted slice which have all element of all slice
 */

func MergeSlicesSorted[T interface{}](comparator types.Comparator, slices ...[]T) []T {
	result := make([]T, 0, defaultCapacity)
	for _, coll := range slices {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return comparator(result[i], result[j]) < 0
	})
	return result
}

/**
 * Create new slice which contains only one element
 * @parameters an item to create slice
 * @returns A slice which contains only one element
 */

func SingletonSlice[T any](item T) []T {
	slice := make([]T, 1, defaultCapacity)
	slice[0] = item
	return slice
}

/**
 * Reverses the order of the given slice.
 * @examples:
 * Reverse([]int{1, 2, 3}) = []int{3, 2, 1}
 *
 * @parameters slice need to reverse
 * @returns A reversed slice
 */

func Reverse[T any](slice []T) {
	size := len(slice)
	for i := 0; i < size/2; i++ {
		swap := slice[i]
		slice[i] = slice[size-i-1]
		slice[size-i-1] = swap
	}
}

/**
 * Get the first element from the given slice
 * @examples:
 * First([]int{}) = nil
 * First([]int{1, 2, 3}) = 1
 * @parameters A slice need to get the first element
 * @returns the first element or nil
 */

func First[T any](slice []T) interface{} {
	if IsEmpty(slice) {
		return nil
	}
	return slice[0]
}

/**
 * Get the last element from the given slice
 * @examples:
 * Last([]int{}) = nil
 * Last([]int{1, 2, 3}) = 3
 *
 * @parameters A slice need to get the first element
 * @returns the last element or nil
 */

func Last[T any](slice []T) interface{} {
	if IsEmpty(slice) {
		return nil
	}
	return slice[len(slice)-1]
}

/**
 * Returns a new slice with only unique elements from the given slice
 * @examples:
 * Unique([]int{1, 2, 3, 2, 4}) = []int{1, 2, 3, 4}
 *
 * @parameters A slice
 * @returns unique slice
 */

func Unique[T comparable](slice []T) []T {
	if IsEmpty(slice) {
		return slice
	}

	result := make([]T, 0, len(slice))
	m := make(map[T]interface{})
	for _, item := range slice {
		if _, ok := m[item]; !ok {
			result = append(result, item)
			m[item] = nil
		}
	}
	return result
}

/**
 * Get index of an @item from given slice
 * @examples:
 * IndexOf([]int{1, 2, 3, 2, 4}, 1) = 0
 * IndexOf([]int{1, 2, 3, 2, 4}, 10) = -1
 * IndexOf([]int{}, 1) = -1
 *
 * @parameters
 * slice - a slice contains data
 * item - need to check index
 * @returns index of item
 */

func IndexOf[T comparable](slice []T, item T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return i
		}
	}
	return indexNotFound

}

/**
 * Check an item is in given slice
 * @examples:
 * Contains([]int{1, 2, 3, 2, 4}, 1) = true
 * Contains([]int{1, 2, 3, 2, 4}, 10) = false
 * Contains([]int{}, 1) = false
 *
 * @parameters
 * slice - a slice contains data
 * item - need to check contains or not
 * @returns  {@code true} if item is in slice
 */

func Contains[T comparable](slice []T, item T) bool {
	return IndexOf(slice, item) != indexNotFound
}

/**
 * Check an item is not in given slice
 * @examples:
 * NotContains([]int{1, 2, 3, 2, 4}, 1) = false
 * NotContains([]int{1, 2, 3, 2, 4}, 10) = true
 * NotContains([]int{}, 1) = true
 *
 * @parameters
 * slice - a slice contains data
 * item - need to check if not contains
 * @returns  {@code true} if item is not in slice
 */

func NotContains[T comparable](slice []T, item T) bool {
	return !Contains(slice, item)
}

/**
 * Returns a slice} containing the intersection of the given slices
 *
 * @parameters two slice to get intersection data
 * @returns intersection slice
 */

func Intersection[T comparable](a, b []T) []T {
	if IsEmpty(a) || IsEmpty(b) {
		return []T{}
	}

	m := make(map[T]bool)
	result := make([]T, 0, defaultCapacity)

	for _, item := range a {
		m[item] = false
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			m[item] = true
		}
	}

	for key, val := range m {
		if val {
			result = append(result, key)
		}
	}
	return result
}

/**
 * Returns a slice containing the union of the given slices
 *
 * @parameters two slice to get union data
 * @returns union slice
 */

func Union[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	result := make([]T, 0, defaultCapacity)

	for _, item := range a {
		m[item] = false
		result = append(result, item)
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

/**
 * Returns a slice containing the difference of the given slices
 *
 * @parameters two slice to get difference data
 * @returns difference slice
 */

func Difference[T comparable](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			m[item] = false
		}
	}
	result := make([]T, 0, defaultCapacity)
	for key, val := range m {
		if val {
			result = append(result, key)
		}
	}

	return result
}

/**
 * Check if two slices are equal
 * @examples:
 * IsEqual([]int{1, 2, 3}, []int{1, 2, 3}) = true
 * IsEqual([]int{1, 2}, []int{1, 3}) = false
 *
 * @parameters two slice to get difference data
 * @returns {@code true} if the string is empty
 */

func IsEqual[T comparable](slice1, slice2 []T) bool {
	size1 := len(slice1)
	size2 := len(slice2)
	if size1 != size2 {
		return false
	}
	for i := 0; i < size1; i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

/**
 * Returns the first element match @predicate
 * @examples:
 * Find([]int{1, 2, 3}, item == 1) = 1
 * Find([]int{1, 2}, item == 3) = nil
 *
 * @parameters
 * slice - given slice data
 * predicate - the checker
 * @returns {@code true} if the string is empty
 */

func Find[T any](slice []T, predicate types.Predicate) interface{} {
	size := len(slice)
	if size == 0 {
		return nil
	}

	for _, item := range slice {
		if predicate(item) {
			return item
		}
	}

	return nil
}

/**
 * Count all elements matched to @predicate
 * @examples:
 * CountMatches([]int{1, 2, 3, 1}, item == 1) = 2
 * CountMatches([]int{1, 2}, item == 3) = 0
 *
 * @parameters
 * slice - given slice data
 * predicate - the checker
 * @returns number of matched item
 */

func CountMatches[T any](slice []T, predicate types.Predicate) int {
	size := len(slice)
	if size == 0 {
		return 0
	}
	matches := 0
	for _, item := range slice {
		if predicate(item) {
			matches++
		}
	}

	return matches
}

/**
 * Remove all elements from given @slice which contains in @remove slice
 * @examples:
 * RemoveAll([]int{1, 2, 3, 1}, []int{1) = []int{2, 3}
 *
 * @parameters
 * slice - given slice data
 * predicate - the checker
 * @returns removed slice
 */

func RemoveAll[T comparable](slice []T, remove []T) []T {
	if IsEmpty(slice) {
		return slice
	}
	removed := make([]T, 0, defaultCapacity)
	for _, item := range slice {
		if NotContains(remove, item) {
			removed = append(removed, item)
		}
	}

	return removed
}
