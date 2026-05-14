package lib

import "math"

type Vector3[T Number] struct {
	X, Y, Z T
}
type Vector3F = Vector3[float32]
type Vector3I = Vector3[int32]

func InitVector3[T Number](x, y, z T) Vector3[T] {
	return Vector3[T]{
		X: x,
		Y: y,
		Z: z,
	}
}
func InitVector3F(x, y, z float32) Vector3F {
	return Vector3F{
		X: x,
		Y: y,
		Z: z,
	}
}

func (self Vector3[T]) Scale(s T) Vector3[T] {
	self.X *= s
	self.Y *= s
	self.Z *= s
	return self
}

func (self Vector3[T]) Divscale(s T) Vector3[T] {
	self.X /= s
	self.Y /= s
	self.Z /= s
	return self
}

func (self Vector3[T]) Offset(x, y, z T) Vector3[T] {
	self.X += x
	self.Y += y
	self.Z += z
	return self
}

func (self Vector3[T]) MulMat(mat Matrix4x4[T]) Vector3[T] {
	return mat.MulVec(self)
}

func (self Vector3[T]) Apply(fn func(x T) (y T)) Vector3[T] {
	self.X = fn(self.X)
	self.Y = fn(self.Y)
	self.Z = fn(self.Z)
	return self
}

func (self Vector3[T]) Min(other Vector3[T]) Vector3[T] {
	self.X = min(self.X, other.X)
	self.Y = min(self.Y, other.Y)
	self.Z = min(self.Z, other.Z)
	return self
}

func (self Vector3[T]) Max(other Vector3[T]) Vector3[T] {
	self.X = max(self.X, other.X)
	self.Y = max(self.Y, other.Y)
	self.Z = max(self.Z, other.Z)
	return self
}

func (self Vector3[T]) Minmax(other ...Vector3[T]) (min Vector3[T], max Vector3[T]) {
	min, max = self, self
	for _, o := range other {
		min = min.Min(o)
		max = max.Max(o)
	}
	return min, max
}

func (self Vector3[T]) Add(other Vector3[T]) Vector3[T] {
	self.X += other.X
	self.Y += other.Y
	self.Z += other.Z
	return self
}

func (self Vector3[T]) Sub(other Vector3[T]) Vector3[T] {
	self.X -= other.X
	self.Y -= other.Y
	self.Z -= other.Z
	return self
}

func (self Vector3[T]) Cross(other Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: self.Y*other.Z - self.Z*other.Y,
		Y: self.Z*other.X - self.X*other.Z,
		Z: self.X*other.Y - self.Y*other.X,
	}
}

func (self Vector3[T]) Dot(other Vector3[T]) T {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self Vector3[T]) Equals(other Vector3[T]) bool {
	return self.X == other.X && self.Y == other.Y && self.Z == other.Z
}

func (self Vector3[T]) Distance(other Vector3[T]) T {
	return self.Sub(other).Length()
}

func (self Vector3[T]) Manhattan(other Vector3[T]) T {
	return self.Sub(other).Abs().Sum()
}

func (self Vector3[T]) Negate() Vector3[T] {
	self.X = -self.X
	self.Y = -self.Y
	self.Z = -self.Z
	return self
}

func (self Vector3[T]) InnerProduct() T {
	return self.Dot(self)
}

func (self Vector3[T]) Length() T {
	return T(math.Sqrt(float64(self.InnerProduct())))
}

func (self Vector3[T]) Volume() T {
	return self.X * self.Y * self.Z
}

func (self Vector3[T]) Sum() T {
	return self.X + self.Y + self.Z
}

func (self Vector3[T]) Normalize() Vector3[T] {
	return self.Divscale(self.Length())
}

func (self Vector3[T]) Abs() Vector3[T] {
	self.X = abs(self.X)
	self.Y = abs(self.Y)
	self.Z = abs(self.Z)
	return self
}

func (self Vector3[T]) LShift32(s int32) Vector3[int32] {
	X := int32(self.X) << s
	Y := int32(self.Y) << s
	Z := int32(self.Z) << s
	return Vector3[int32]{X: X, Y: Y, Z: Z}
}

func (self Vector3[T]) LShift64(s int64) Vector3[int64] {
	X := int64(self.X) << s
	Y := int64(self.Y) << s
	Z := int64(self.Z) << s
	return Vector3[int64]{X: X, Y: Y, Z: Z}
}

func (self Vector3[T]) RShift32(s int32) Vector3[int32] {
	X := int32(self.X) >> s
	Y := int32(self.Y) >> s
	Z := int32(self.Z) >> s
	return Vector3[int32]{X: X, Y: Y, Z: Z}
}

func (self Vector3[T]) RShift64(s int64) Vector3[int64] {
	X := int64(self.X) >> s
	Y := int64(self.Y) >> s
	Z := int64(self.Z) >> s
	return Vector3[int64]{X: X, Y: Y, Z: Z}
}

func FloorToInt[T Number, I Integer](self Vector3[T]) Vector3[I] {
	x := I(math.Floor(float64(self.X)))
	y := I(math.Floor(float64(self.Y)))
	z := I(math.Floor(float64(self.Z)))
	return Vector3[I]{X: x, Y: y, Z: z}
}

func CeilToInt[T Number, I Integer](self Vector3[T]) Vector3[I] {
	x := I(math.Ceil(float64(self.X)))
	y := I(math.Ceil(float64(self.Y)))
	z := I(math.Ceil(float64(self.Z)))
	return Vector3[I]{X: x, Y: y, Z: z}
}

func RoundToInt[T Number, I Integer](self Vector3[T]) Vector3[I] {
	x := I(math.Round(float64(self.X)))
	y := I(math.Round(float64(self.Y)))
	z := I(math.Round(float64(self.Z)))
	return Vector3[I]{X: x, Y: y, Z: z}
}

func (self Vector3[T]) FloorToInt32() Vector3[int32] {
	return FloorToInt[T, int32](self)
}

func (self Vector3[T]) FloorToInt64() Vector3[int64] {
	return FloorToInt[T, int64](self)
}

func (self Vector3[T]) CeilToInt32() Vector3[int32] {
	return CeilToInt[T, int32](self)
}

func (self Vector3[T]) CeilToInt64() Vector3[int64] {
	return CeilToInt[T, int64](self)
}

func (self Vector3[T]) RoundToInt32() Vector3[int32] {
	return RoundToInt[T, int32](self)
}

func (self Vector3[T]) RoundToInt64() Vector3[int64] {
	return RoundToInt[T, int64](self)
}

func (self Vector3[T]) ToFloat32() Vector3[float32] {
	return Vector3[float32]{
		float32(self.X),
		float32(self.Y),
		float32(self.Z),
	}
}

func (self Vector3[T]) ToFloat64() Vector3[float64] {
	return Vector3[float64]{
		float64(self.X),
		float64(self.Y),
		float64(self.Z),
	}
}

func (self Vector3[T]) ToFixedSlice() [3]T {
	return [3]T{self.X, self.Y, self.Z}
}

func (self Vector3[T]) ToSlice() []T {
	return []T{self.X, self.Y, self.Z}
}

func (self Vector3[T]) SetX(x T) Vector3[T] {
	self.X = x
	return self
}
func (self Vector3[T]) SetY(y T) Vector3[T] {
	self.Y = y
	return self
}
func (self Vector3[T]) SetZ(z T) Vector3[T] {
	self.Z = z
	return self
}
func (self Vector3[T]) Set(x, y, z T) Vector3[T] {
	return Vector3[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

func (self Vector3[T]) FromFixedSlice(slice [3]T) Vector3[T] {
	return Vector3[T]{
		X: slice[0],
		Y: slice[1],
		Z: slice[2],
	}
}
func (self Vector3[T]) FromSlice(slice []T) Vector3[T] {
	return Vector3[T]{
		X: slice[0],
		Y: slice[1],
		Z: slice[2],
	}
}
