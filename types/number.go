package types

type IntNumber interface {
	int | int8 | int16 | int32 | int64
}

type UnitNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type FloatNumber interface {
	float32 | float64
}

type Number interface {
	IntNumber | FloatNumber | UnitNumber
}
