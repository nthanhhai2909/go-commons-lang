package intutils

type IntNumber interface {
	int | int8 | int16 | int32 | int64
}

func Abs[T IntNumber](num T) T {
	if num > 0 {
		return num
	}
	return -num
}

func Sum[T IntNumber](nums ...T) T {
	var sum T
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max[T IntNumber](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T IntNumber](a, b T) T {
	if a < b {
		return a
	}
	return b
}
