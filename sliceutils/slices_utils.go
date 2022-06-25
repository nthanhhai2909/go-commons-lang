package sliceutils

import (
	"sort"
)

const (
	DefaultCapacity = 10
	indexNotFound   = -1
)

func IsEmpty[T any](slice []T) bool {
	if len(slice) == 0 {
		return true
	}
	return false
}

func IsNotEmpty[T any](slice []T) bool {
	return !IsEmpty(slice)
}

func MergeSlices[T any](slices ...[]T) []T {
	result := make([]T, 0, DefaultCapacity)
	for _, coll := range slices {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	return result
}

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

func SingletonSlice[T any](item T) []T {
	slice := make([]T, 1, DefaultCapacity)
	slice[0] = item
	return slice
}

func Reverse[T any](slice []T) {
	size := len(slice)
	for i := 0; i < size/2; i++ {
		swap := slice[i]
		slice[i] = slice[size-i-1]
		slice[size-i-1] = swap
	}
}

func First[T any](slice []T) interface{} {
	if IsEmpty(slice) {
		return nil
	}
	return slice[0]
}

func Last[T any](slice []T) interface{} {
	if IsEmpty(slice) {
		return nil
	}
	return slice[len(slice)-1]
}

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

func IndexOf[T comparable](slice []T, item T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return i
		}
	}
	return indexNotFound

}

func Contains[T comparable](slice []T, item T) bool {
	return IndexOf(slice, item) != indexNotFound
}

func NotContains[T comparable](slice []T, item T) bool {
	return !Contains(slice, item)
}
