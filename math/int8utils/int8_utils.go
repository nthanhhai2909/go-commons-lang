package int8utils

func Abs(num int8) int8 {
	if num > 0 {
		return num
	}
	return -num
}

func Sum(nums ...int8) int8 {
	var sum int8
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}
