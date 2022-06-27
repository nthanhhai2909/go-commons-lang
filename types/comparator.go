package types

import (
	"strings"
	"time"
)

type Comparator func(a, b interface{}) int

func StringComparator(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

func InverseStringComparator(a, b interface{}) int {
	return -StringComparator(a, b)
}

func IntComparator(a, b interface{}) int {
	return NumberComparator(a.(int), b.(int))
}

func InverseIntComparator(a, b interface{}) int {
	return -IntComparator(a, b)
}

func Int8Comparator(a, b interface{}) int {
	return NumberComparator(a.(int8), b.(int8))
}

func InverseInt8Comparator(a, b interface{}) int {
	return -Int8Comparator(a, b)
}

func Int16Comparator(a, b interface{}) int {
	return NumberComparator(a.(int16), b.(int16))
}

func InverseInt16Comparator(a, b interface{}) int {
	return -Int16Comparator(a, b)
}

func Int32Comparator(a, b interface{}) int {
	return NumberComparator(a.(int32), b.(int32))
}

func InverseInt32Comparator(a, b interface{}) int {
	return -Int32Comparator(a, b)
}

func Int64Comparator(a, b interface{}) int {
	return NumberComparator(a.(int64), b.(int64))
}

func InverseInt64Comparator(a, b interface{}) int {
	return -Int64Comparator(a, b)
}

func UintComparator(a, b interface{}) int {
	return NumberComparator(a.(uint), b.(uint))
}

func InverseUintComparator(a, b interface{}) int {
	return -UintComparator(a, b)
}

func Uint8Comparator(a, b interface{}) int {
	return NumberComparator(a.(uint8), b.(uint8))
}

func InverseUint8Comparator(a, b interface{}) int {
	return -Uint8Comparator(a, b)
}

func Uint16Comparator(a, b interface{}) int {
	return NumberComparator(a.(uint16), b.(uint16))
}

func InverseUint16Comparator(a, b interface{}) int {
	return -Uint16Comparator(a, b)
}

func Uint32Comparator(a, b interface{}) int {
	return NumberComparator(a.(uint32), b.(uint32))
}

func InverseUint32Comparator(a, b interface{}) int {
	return -Uint32Comparator(a, b)
}

func Uint64Comparator(a, b interface{}) int {
	return NumberComparator(a.(uint64), b.(uint64))
}

func InverseUint64Comparator(a, b interface{}) int {
	return -Uint64Comparator(a, b)
}

func Float32Comparator(a, b interface{}) int {
	return NumberComparator(a.(float32), b.(float32))
}

func InverseFloat32Comparator(a, b interface{}) int {
	return -Float32Comparator(a, b)
}

func Float64Comparator(a, b interface{}) int {
	return NumberComparator(a.(float64), b.(float64))
}

func InverseFloat64Comparator(a, b interface{}) int {
	return -Float64Comparator(a, b)
}

func ByteComparator(a, b interface{}) int {
	return NumberComparator(a.(byte), b.(byte))
}

func InverseByteComparator(a, b interface{}) int {
	return -ByteComparator(a, b)
}

func TimeComparator(a, b interface{}) int {
	aAsserted := a.(time.Time)
	bAsserted := b.(time.Time)

	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}

func InverseTimeComparator(a, b interface{}) int {
	return -TimeComparator(a, b)
}

func NumberComparator[T Number](a, b T) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}
