package other

import (
	"math"
)

type Integer interface {
	SInteger | UInteger
}
type SInteger interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}
type UInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}
type Number interface {
	~float32 | ~float64 | Integer
}

func Sincos[T Number](theta T) (s T, c T) {
	s64, c64 := math.Sincos(float64(theta))
	return T(s64), T(c64)
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Floor[T Number](x T) T {
	return T(math.Floor(float64(x)))
}

func Ceil[T Number](x T) T {
	return T(math.Ceil(float64(x)))
}

func Round[T Number](x T) T {
	return T(math.Round(float64(x)))
}

func Mod[T Number](x T, m T) T {
	return T(math.Mod(float64(x), float64(m)))
}

func Tan[T Number](x T) T {
	return T(math.Tan(float64(x)))
}
