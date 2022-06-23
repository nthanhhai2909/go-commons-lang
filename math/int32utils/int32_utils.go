package int8utils

func Abs(num int32) int32 {
	if num > 0 {
		return num
	}
	return -num
}

func Sum(nums ...int32) int32 {
	var sum int32
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
