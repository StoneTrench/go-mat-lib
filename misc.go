package lib

import (
	"errors"
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

func sincos[T Number](theta T) (s T, c T) {
	s64, c64 := math.Sincos(float64(theta))
	return T(s64), T(c64)
}

func abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func tan[T Number](x T) T {
	return T(math.Tan(float64(x)))
}

func err_append(msg string, err error) error {
	return errors.Join(errors.New(msg), err)
}
