package int8utils

func Abs(num int16) int16 {
	if num > 0 {
		return num
	}
	return -num
}

func Sum(nums ...int16) int16 {
	var sum int16
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}
