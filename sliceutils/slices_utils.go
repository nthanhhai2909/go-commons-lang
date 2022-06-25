package sliceutils

import (
	"sort"
)

const (
	DefaultCapacity = 10
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
	result := make([]T, 0, DefaultCapacity)
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

func MergeSlicesSorted[T any](comparator func(a, b T) bool, slices ...[]T) []T {
	result := make([]T, 0, DefaultCapacity)
	for _, coll := range slices {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return comparator(result[i], result[j])
	})
	return result
}

/**
 * Create new slice which contains only one element
 * @parameters an item to create slice
 * @returns A slice which contains only one element
 */

func SingletonSlice[T any](item T) []T {
	slice := make([]T, 1, DefaultCapacity)
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
 * @returns  {@code true} if item is not in slice
 */

func NotContains[T comparable](slice []T, item T) bool {
	return !Contains(slice, item)
}
