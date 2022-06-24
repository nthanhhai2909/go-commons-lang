package sliceutils

import (
	"sort"
)

const (
	DefaultCapacity = 10
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

func First[T any](slice []T) (interface{}, bool) {
	if IsEmpty(slice) {
		return nil, false
	}
	return slice[0], true
}

func Last[T any](slice []T) (interface{}, bool) {
	if IsEmpty(slice) {
		return nil, false
	}
	return slice[len(slice)-1], true
}
