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
