package types

import (
	"testing"
	"time"
)

func TestStringComparator(t *testing.T) {
	if StringComparator("a", "a") != 0 ||
		StringComparator("a", "b") != -1 ||
		StringComparator("b", "a") != 1 {
		t.Errorf("StringComparator failed")
	}
}

func TestInverseStringComparator(t *testing.T) {
	if InverseStringComparator("a", "a") != 0 ||
		InverseStringComparator("a", "b") != 1 ||
		InverseStringComparator("b", "a") != -1 {
		t.Errorf("InverseStringComparator failed")
	}
}

func TestIntComparator(t *testing.T) {
	if IntComparator(1, 1) != 0 ||
		IntComparator(1, 2) != -1 ||
		IntComparator(2, -1) != 1 {
		t.Errorf("StringComparator failed")
	}
}

func TestInverseIntComparator(t *testing.T) {
	if InverseIntComparator(1, 1) != 0 ||
		InverseIntComparator(1, 2) != 1 ||
		InverseIntComparator(2, -1) != -1 {
		t.Errorf("InverseIntComparator failed")
	}
}

func TestInt8Comparator(t *testing.T) {
	if Int8Comparator(int8(1), int8(1)) != 0 ||
		Int8Comparator(int8(1), int8(2)) != -1 ||
		Int8Comparator(int8(2), int8(-128)) != 1 {
		t.Errorf("Int8Comparator failed")
	}
}

func TestInverseInt8Comparator(t *testing.T) {
	if InverseInt8Comparator(int8(1), int8(1)) != 0 ||
		InverseInt8Comparator(int8(1), int8(2)) != 1 ||
		InverseInt8Comparator(int8(2), int8(-128)) != -1 {
		t.Errorf("Int8Comparator failed")
	}
}

func TestInt16Comparator(t *testing.T) {
	if Int16Comparator(int16(1), int16(1)) != 0 ||
		Int16Comparator(int16(1), int16(2)) != -1 ||
		Int16Comparator(int16(2), int16(-128)) != 1 {
		t.Errorf("Int16Comparator failed")
	}
}

func TestInverseInt16Comparator(t *testing.T) {
	if InverseInt16Comparator(int16(1), int16(1)) != 0 ||
		InverseInt16Comparator(int16(1), int16(2)) != 1 ||
		InverseInt16Comparator(int16(2), int16(-128)) != -1 {
		t.Errorf("InverseInt16Comparator failed")
	}
}

func TestInt32Comparator(t *testing.T) {
	if Int32Comparator(int32(1), int32(1)) != 0 ||
		Int32Comparator(int32(1), int32(2)) != -1 ||
		Int32Comparator(int32(2), int32(-128)) != 1 {
		t.Errorf("Int32Comparator failed")
	}
}

func TestInverseInt32Comparator(t *testing.T) {
	if InverseInt32Comparator(int32(1), int32(1)) != 0 ||
		InverseInt32Comparator(int32(1), int32(2)) != 1 ||
		InverseInt32Comparator(int32(2), int32(-128)) != -1 {
		t.Errorf("InverseInt32Comparator failed")
	}
}

func TestInt64Comparator(t *testing.T) {
	if Int64Comparator(int64(1), int64(1)) != 0 ||
		Int64Comparator(int64(1), int64(2)) != -1 ||
		Int64Comparator(int64(2), int64(-128)) != 1 {
		t.Errorf("InverseInt64Comparator failed")
	}
}

func TestInverseInt64Comparator(t *testing.T) {
	if InverseInt64Comparator(int64(1), int64(1)) != 0 ||
		InverseInt64Comparator(int64(1), int64(2)) != 1 ||
		InverseInt64Comparator(int64(2), int64(-128)) != -1 {
		t.Errorf("InverseInt64Comparator failed")
	}
}

func TestUintComparator(t *testing.T) {
	if UintComparator(uint(1), uint(1)) != 0 ||
		UintComparator(uint(1), uint(2)) != -1 ||
		UintComparator(uint(2), uint(1)) != 1 {
		t.Errorf("UintComparator failed")
	}
}

func TestInverseUintComparator(t *testing.T) {
	if InverseUintComparator(uint(1), uint(1)) != 0 ||
		InverseUintComparator(uint(1), uint(2)) != 1 ||
		InverseUintComparator(uint(2), uint(1)) != -1 {
		t.Errorf("InverseUintComparator failed")
	}
}

func TestUint8Comparator(t *testing.T) {
	if Uint8Comparator(uint8(1), uint8(1)) != 0 ||
		Uint8Comparator(uint8(1), uint8(2)) != -1 ||
		Uint8Comparator(uint8(2), uint8(1)) != 1 {
		t.Errorf("Uint8Comparator failed")
	}
}

func TestInverseUint8Comparator(t *testing.T) {
	if InverseUint8Comparator(uint8(1), uint8(1)) != 0 ||
		InverseUint8Comparator(uint8(1), uint8(2)) != 1 ||
		InverseUint8Comparator(uint8(2), uint8(1)) != -1 {
		t.Errorf("InverseUint8Comparator failed")
	}
}

func TestUint16Comparator(t *testing.T) {
	if Uint16Comparator(uint16(1), uint16(1)) != 0 ||
		Uint16Comparator(uint16(1), uint16(2)) != -1 ||
		Uint16Comparator(uint16(2), uint16(1)) != 1 {
		t.Errorf("Uint16Comparator failed")
	}
}

func TestInverseUint16Comparator(t *testing.T) {
	if InverseUint16Comparator(uint16(1), uint16(1)) != 0 ||
		InverseUint16Comparator(uint16(1), uint16(2)) != 1 ||
		InverseUint16Comparator(uint16(2), uint16(1)) != -1 {
		t.Errorf("InverseUint16Comparator failed")
	}
}

func TestUint32Comparator(t *testing.T) {
	if Uint32Comparator(uint32(1), uint32(1)) != 0 ||
		Uint32Comparator(uint32(1), uint32(2)) != -1 ||
		Uint32Comparator(uint32(2), uint32(1)) != 1 {
		t.Errorf("Uint32Comparator failed")
	}
}

func TestInverseUint32Comparator(t *testing.T) {
	if InverseUint32Comparator(uint32(1), uint32(1)) != 0 ||
		InverseUint32Comparator(uint32(1), uint32(2)) != 1 ||
		InverseUint32Comparator(uint32(2), uint32(1)) != -1 {
		t.Errorf("InverseUint32Comparator failed")
	}
}

func TestUint64Comparator(t *testing.T) {
	if Uint64Comparator(uint64(1), uint64(1)) != 0 ||
		Uint64Comparator(uint64(1), uint64(2)) != -1 ||
		Uint64Comparator(uint64(2), uint64(1)) != 1 {
		t.Errorf("Uint64Comparator failed")
	}
}

func TestInverseUint64Comparator(t *testing.T) {
	if InverseUint64Comparator(uint64(1), uint64(1)) != 0 ||
		InverseUint64Comparator(uint64(1), uint64(2)) != 1 ||
		InverseUint64Comparator(uint64(2), uint64(1)) != -1 {
		t.Errorf("InverseUint64Comparator failed")
	}
}

func TestFloat32Comparator(t *testing.T) {
	if Float32Comparator(float32(1), float32(1)) != 0 ||
		Float32Comparator(float32(1), float32(2)) != -1 ||
		Float32Comparator(float32(2), float32(1)) != 1 {
		t.Errorf("Float32Comparator failed")
	}
}

func TestInverseFloat32Comparator(t *testing.T) {
	if InverseFloat32Comparator(float32(1), float32(1)) != 0 ||
		InverseFloat32Comparator(float32(1), float32(2)) != 1 ||
		InverseFloat32Comparator(float32(2), float32(1)) != -1 {
		t.Errorf("InverseFloat32Comparator failed")
	}
}

func TestFloat64Comparator(t *testing.T) {
	if Float64Comparator(float64(1), float64(1)) != 0 ||
		Float64Comparator(float64(1), float64(2)) != -1 ||
		Float64Comparator(float64(2), float64(1)) != 1 {
		t.Errorf("Float64Comparator failed")
	}
}

func TestInverseFloat64Comparator(t *testing.T) {
	if InverseFloat64Comparator(float64(1), float64(1)) != 0 ||
		InverseFloat64Comparator(float64(1), float64(2)) != 1 ||
		InverseFloat64Comparator(float64(2), float64(1)) != -1 {
		t.Errorf("InverseFloat64Comparator failed")
	}
}

func TestByteComparator(t *testing.T) {
	if ByteComparator(byte(1), byte(1)) != 0 ||
		ByteComparator(byte(1), byte(2)) != -1 ||
		ByteComparator(byte(2), byte(1)) != 1 {
		t.Errorf("ByteComparator failed")
	}
}

func TestInverseByteComparator(t *testing.T) {
	if InverseByteComparator(byte(1), byte(1)) != 0 ||
		InverseByteComparator(byte(1), byte(2)) != 1 ||
		InverseByteComparator(byte(2), byte(1)) != -1 {
		t.Errorf("InverseByteComparator failed")
	}
}

func TestTimeComparator(t *testing.T) {
	time1 := time.Now()
	time2 := time1
	time3 := time1.Add(time.Hour)
	if TimeComparator(time1, time2) != 0 ||
		TimeComparator(time1, time3) != -1 ||
		TimeComparator(time3, time1) != 1 {
		t.Errorf("TimeComparator failed")
	}
}

func TestInverseTimeComparator(t *testing.T) {
	time1 := time.Now()
	time2 := time1
	time3 := time1.Add(time.Hour)
	if InverseTimeComparator(time1, time2) != 0 ||
		InverseTimeComparator(time1, time3) != 1 ||
		InverseTimeComparator(time3, time1) != -1 {
		t.Errorf("InverseTimeComparator failed")
	}
}
