package main

type Integer interface {
	int | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Numeric interface {
	Float | Integer
}

func min[T Integer | Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func max[T Numeric](a, b T) T {
	if a < b {
		return b
	}
	return a
}
