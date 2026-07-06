package vec

import (
	. "github.com/StoneTrench/go-mat-lib/const"
)

type Matrix4x4[T Number] struct {
	M11, M12, M13, M14,
	M21, M22, M23, M24,
	M31, M32, M33, M34,
	M41, M42, M43, M44 T
}

func (mat Matrix4x4[T]) InitTranslation(x, y, z T) Matrix4x4[T] {
	return Matrix4x4[T]{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}
func (mat Matrix4x4[T]) InitIdentity(s T) Matrix4x4[T] {
	return mat.InitDiagonal(s, s, s, s)
}
func (mat Matrix4x4[T]) InitDiagonal(a, b, c, d T) Matrix4x4[T] {
	return Matrix4x4[T]{
		a, 0, 0, 0,
		0, b, 0, 0,
		0, 0, c, 0,
		0, 0, 0, d,
	}
}
func (mat Matrix4x4[T]) InitRotX(theta T) Matrix4x4[T] {
	s, c := Sincos(theta)

	return Matrix4x4[T]{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}
func (mat Matrix4x4[T]) InitRotY(theta T) Matrix4x4[T] {
	s, c := Sincos(theta)

	return Matrix4x4[T]{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}
func (mat Matrix4x4[T]) InitRotZ(theta T) Matrix4x4[T] {
	s, c := Sincos(theta)

	return Matrix4x4[T]{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}
func (mat Matrix4x4[T]) InitPerspectiveProjection(fovy, width, height, near, far T) Matrix4x4[T] {
	tanHalfFovy := Tan(fovy / 2)
	aspect := width / height

	m00 := 1 / (aspect * tanHalfFovy)
	m11 := -(1 / tanHalfFovy)

	m22 := far / (near - far)
	m23 := -T(1)
	m32 := (far * near) / (near - far)

	return Matrix4x4[T]{
		m00, 0, 0, 0,
		0, m11, 0, 0,
		0, 0, m22, m23,
		0, 0, m32, 0,
	}
}
func (mat Matrix4x4[T]) InitLookAt(eye, center, up Vector3[T]) Matrix4x4[T] {
	f := center.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)

	return Matrix4x4[T]{
		s.X, u.X, -f.X, 0,
		s.Y, u.Y, -f.Y, 0,
		s.Z, u.Z, -f.Z, 0,
		-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1,
	}
}

func (mat Matrix4x4[T]) RotateX(theta T) Matrix4x4[T] {
	return mat.MulMat(mat.InitRotX(theta))
}
func (mat Matrix4x4[T]) RotateY(theta T) Matrix4x4[T] {
	return mat.MulMat(mat.InitRotY(theta))
}
func (mat Matrix4x4[T]) RotateZ(theta T) Matrix4x4[T] {
	return mat.MulMat(mat.InitRotZ(theta))
}

func (mat Matrix4x4[T]) MulVec(vec Vector3[T]) Vector3[T] {
	return Vector3[T]{
		vec.X*mat.M11 + vec.Y*mat.M12 + vec.Z*mat.M13 + mat.M14,
		vec.X*mat.M21 + vec.Y*mat.M22 + vec.Z*mat.M23 + mat.M24,
		vec.X*mat.M31 + vec.Y*mat.M32 + vec.Z*mat.M33 + mat.M34,
	}
}
func (mat Matrix4x4[T]) MulMat(left Matrix4x4[T]) Matrix4x4[T] {
	return Matrix4x4[T]{
		M11: mat.M11*left.M11 + mat.M12*left.M21 + mat.M13*left.M31 + mat.M14*left.M41,
		M12: mat.M11*left.M12 + mat.M12*left.M22 + mat.M13*left.M32 + mat.M14*left.M42,
		M13: mat.M11*left.M13 + mat.M12*left.M23 + mat.M13*left.M33 + mat.M14*left.M43,
		M14: mat.M11*left.M14 + mat.M12*left.M24 + mat.M13*left.M34 + mat.M14*left.M44,

		M21: mat.M21*left.M11 + mat.M22*left.M21 + mat.M23*left.M31 + mat.M24*left.M41,
		M22: mat.M21*left.M12 + mat.M22*left.M22 + mat.M23*left.M32 + mat.M24*left.M42,
		M23: mat.M21*left.M13 + mat.M22*left.M23 + mat.M23*left.M33 + mat.M24*left.M43,
		M24: mat.M21*left.M14 + mat.M22*left.M24 + mat.M23*left.M34 + mat.M24*left.M44,

		M31: mat.M31*left.M11 + mat.M32*left.M21 + mat.M33*left.M31 + mat.M34*left.M41,
		M32: mat.M31*left.M12 + mat.M32*left.M22 + mat.M33*left.M32 + mat.M34*left.M42,
		M33: mat.M31*left.M13 + mat.M32*left.M23 + mat.M33*left.M33 + mat.M34*left.M43,
		M34: mat.M31*left.M14 + mat.M32*left.M24 + mat.M33*left.M34 + mat.M34*left.M44,

		M41: mat.M41*left.M11 + mat.M42*left.M21 + mat.M43*left.M31 + mat.M44*left.M41,
		M42: mat.M41*left.M12 + mat.M42*left.M22 + mat.M43*left.M32 + mat.M44*left.M42,
		M43: mat.M41*left.M13 + mat.M42*left.M23 + mat.M43*left.M33 + mat.M44*left.M43,
		M44: mat.M41*left.M14 + mat.M42*left.M24 + mat.M43*left.M34 + mat.M44*left.M44,
	}
}

func (mat Matrix4x4[T]) Transpose() Matrix4x4[T] {
	return Matrix4x4[T]{
		mat.M11, mat.M21, mat.M31, mat.M41,
		mat.M12, mat.M22, mat.M32, mat.M42,
		mat.M13, mat.M23, mat.M33, mat.M43,
		mat.M14, mat.M24, mat.M34, mat.M44,
	}
}
func (mat Matrix4x4[T]) HTMInverse() Matrix4x4[T] {
	m11, m12, m13 := mat.M11, mat.M21, mat.M31
	m21, m22, m23 := mat.M12, mat.M22, mat.M32
	m31, m32, m33 := mat.M13, mat.M23, mat.M33

	m14 := -(m11*mat.M14 + m12*mat.M24 + m13*mat.M34)
	m24 := -(m21*mat.M14 + m22*mat.M24 + m23*mat.M34)
	m34 := -(m31*mat.M14 + m32*mat.M24 + m33*mat.M34)

	return Matrix4x4[T]{
		m11, m12, m13, m14,
		m21, m22, m23, m24,
		m31, m32, m33, m34,
		0, 0, 0, 1,
	}
}
