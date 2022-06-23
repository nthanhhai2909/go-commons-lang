package int8utils

func Abs(num int64) int64 {
	if num > 0 {
		return num
	}
	return -num
}

func Sum(nums ...int64) int64 {
	var sum int64
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
