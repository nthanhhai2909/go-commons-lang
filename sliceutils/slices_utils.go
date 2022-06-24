package sliceutils

import (
	"sort"
)

func IsEmpty[T any](collection []T) bool {
	if len(collection) == 0 {
		return true
	}
	return false
}

func IsNotEmpty[T any](collection []T) bool {
	return !IsEmpty(collection)
}

func MergeSlices[T any](collections ...[]T) []T {
	result := make([]T, 0)
	for _, coll := range collections {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	return result
}

func MergeSlicesSorted[T any](sortFunc func(a, b T) bool, collections ...[]T) []T {
	result := make([]T, 0)
	for _, coll := range collections {
		for _, item := range coll {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return sortFunc(result[i], result[j])
	})
	return result
}
