package types

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}
