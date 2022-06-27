package intutils

import "go-commons-lang/types"

func Abs[T types.Number](num T) T {
	if num > 0 {
		return num
	}
	return -num
}

func Sum[T types.Number](nums ...T) T {
	var sum T
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max[T types.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T types.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
