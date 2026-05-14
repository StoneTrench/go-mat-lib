package lib

import "math"

type Quaternion[T Number] struct {
	W, X, Y, Z T
}

func InitQuaternion[T Number]() Quaternion[T] {
	return Quaternion[T]{
		W: 1,
		X: 0,
		Y: 0,
		Z: 0,
	}
}
func InitAngleAxis[T Number](angle, x, y, z T) Quaternion[T] {
	sin, cos := sincos(angle / 2)
	return Quaternion[T]{
		W: cos,
		X: x * sin,
		Y: y * sin,
		Z: z * sin,
	}
}
func InitWXYZ[T Number](w, x, y, z T) Quaternion[T] {
	return Quaternion[T]{
		W: w,
		X: x,
		Y: y,
		Z: z,
	}
}

func (q Quaternion[T]) Dot(o Quaternion[T]) T {
	return q.W*o.W + q.X*o.X + q.Y*o.Y + q.Z*o.Z
}
func (q Quaternion[T]) MulQuat(o Quaternion[T]) Quaternion[T] {
	return Quaternion[T]{
		W: q.W*o.W - q.X*o.X - q.Y*o.Y - q.Z*o.Z,
		X: q.W*o.X - q.X*o.Y - q.Y*o.Z - q.Z*o.W,
		Y: q.W*o.Y - q.X*o.Z - q.Y*o.W - q.Z*o.X,
		Z: q.W*o.Z - q.X*o.W - q.Y*o.X - q.Z*o.Y,
	}
}

func (q Quaternion[T]) Normalize() Quaternion[T] {
	len := T(math.Sqrt(float64(q.Dot(q))))
	return Quaternion[T]{
		W: q.W / len,
		X: q.X / len,
		Y: q.Y / len,
		Z: q.Z / len,
	}
}
func (q Quaternion[T]) Conjugate() Quaternion[T] {
	return Quaternion[T]{
		W: q.W,
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
	}
}

// MulVec is a passive rotation.
//
// Active rotation is when the point is rotated with respect to the coordinate system, and passive rotation is when the coordinate system is rotated with respect to the point. The two rotations are opposite from each other.
// func (q Quaternion[T]) MulVec(v Vector3) Vector3 {
// 	vQuat := Quaternion[T]{W: 0, X: v.X, Y: v.Y, Z: v.Z}
// 	res := q.MulQuat(vQuat).MulQuat(q.Conjugate())
// 	return Vector3{X: res.X, Y: res.Y, Z: res.Z}
// }
func (q Quaternion[T]) MulVec(v Vector3[T]) Vector3[T] {
	numX := q.X * 2
	numY := q.Y * 2
	numZ := q.Z * 2

	tx := numY*v.Z - numZ*v.Y
	ty := numZ*v.X - numX*v.Z
	tz := numX*v.Y - numY*v.X

	return Vector3[T]{
		X: v.X + q.W*tx + (q.Y*tz - q.Z*ty),
		Y: v.Y + q.W*ty + (q.Z*tx - q.X*tz),
		Z: v.Z + q.W*tz + (q.X*ty - q.Y*tx),
	}
}
