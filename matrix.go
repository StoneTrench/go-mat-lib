package lib

type Matrix4x4[T Number] struct {
	M11, M12, M13, M14,
	M21, M22, M23, M24,
	M31, M32, M33, M34,
	M41, M42, M43, M44 T
}

func InitTranslation[T Number](x, y, z T) Matrix4x4[T] {
	return Matrix4x4[T]{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}
func InitIdentity[T Number](s T) Matrix4x4[T] {
	return InitDiagonal(s, s, s, s)
}
func InitDiagonal[T Number](a, b, c, d T) Matrix4x4[T] {
	return Matrix4x4[T]{
		a, 0, 0, 0,
		0, b, 0, 0,
		0, 0, c, 0,
		0, 0, 0, d,
	}
}

func InitRotX[T Number](theta T) Matrix4x4[T] {
	s, c := sincos(theta)

	return Matrix4x4[T]{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}
func InitRotY[T Number](theta T) Matrix4x4[T] {
	s, c := sincos(theta)

	return Matrix4x4[T]{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}
func InitRotZ[T Number](theta T) Matrix4x4[T] {
	s, c := sincos(theta)

	return Matrix4x4[T]{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func InitPerspective[T Number](fovy, width, height, near, far T) Matrix4x4[T] {
	tanHalfFovy := tan(fovy / 2)
	aspect := width / height

	m00 := 1 / (aspect * tanHalfFovy)
	m11 := 1 / tanHalfFovy

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

func (self Matrix4x4[T]) RotateX(theta T) Matrix4x4[T] {
	return self.MulMat(InitRotX(theta))
}
func (self Matrix4x4[T]) RotateY(theta T) Matrix4x4[T] {
	return self.MulMat(InitRotY(theta))
}
func (self Matrix4x4[T]) RotateZ(theta T) Matrix4x4[T] {
	return self.MulMat(InitRotZ(theta))
}

func (self Matrix4x4[T]) MulVec(vec Vector3[T]) Vector3[T] {
	return Vector3[T]{
		vec.X*self.M11 + vec.Y*self.M12 + vec.Z*self.M13 + self.M14,
		vec.X*self.M21 + vec.Y*self.M22 + vec.Z*self.M23 + self.M24,
		vec.X*self.M31 + vec.Y*self.M32 + vec.Z*self.M33 + self.M34,
	}
}
func (self Matrix4x4[T]) MulMat(left Matrix4x4[T]) Matrix4x4[T] {
	return Matrix4x4[T]{
		M11: self.M11*left.M11 + self.M12*left.M21 + self.M13*left.M31 + self.M14*left.M41,
		M12: self.M11*left.M12 + self.M12*left.M22 + self.M13*left.M32 + self.M14*left.M42,
		M13: self.M11*left.M13 + self.M12*left.M23 + self.M13*left.M33 + self.M14*left.M43,
		M14: self.M11*left.M14 + self.M12*left.M24 + self.M13*left.M34 + self.M14*left.M44,

		M21: self.M21*left.M11 + self.M22*left.M21 + self.M23*left.M31 + self.M24*left.M41,
		M22: self.M21*left.M12 + self.M22*left.M22 + self.M23*left.M32 + self.M24*left.M42,
		M23: self.M21*left.M13 + self.M22*left.M23 + self.M23*left.M33 + self.M24*left.M43,
		M24: self.M21*left.M14 + self.M22*left.M24 + self.M23*left.M34 + self.M24*left.M44,

		M31: self.M31*left.M11 + self.M32*left.M21 + self.M33*left.M31 + self.M34*left.M41,
		M32: self.M31*left.M12 + self.M32*left.M22 + self.M33*left.M32 + self.M34*left.M42,
		M33: self.M31*left.M13 + self.M32*left.M23 + self.M33*left.M33 + self.M34*left.M43,
		M34: self.M31*left.M14 + self.M32*left.M24 + self.M33*left.M34 + self.M34*left.M44,

		M41: self.M41*left.M11 + self.M42*left.M21 + self.M43*left.M31 + self.M44*left.M41,
		M42: self.M41*left.M12 + self.M42*left.M22 + self.M43*left.M32 + self.M44*left.M42,
		M43: self.M41*left.M13 + self.M42*left.M23 + self.M43*left.M33 + self.M44*left.M43,
		M44: self.M41*left.M14 + self.M42*left.M24 + self.M43*left.M34 + self.M44*left.M44,
	}
}

func (self Matrix4x4[T]) Transpose() Matrix4x4[T] {
	return Matrix4x4[T]{
		self.M11, self.M21, self.M31, self.M41,
		self.M12, self.M22, self.M32, self.M42,
		self.M13, self.M23, self.M33, self.M43,
		self.M14, self.M24, self.M34, self.M44,
	}
}
func (self Matrix4x4[T]) HTMInverse() Matrix4x4[T] {
	m11, m12, m13 := self.M11, self.M21, self.M31
	m21, m22, m23 := self.M12, self.M22, self.M32
	m31, m32, m33 := self.M13, self.M23, self.M33

	m14 := -(m11*self.M14 + m12*self.M24 + m13*self.M34)
	m24 := -(m21*self.M14 + m22*self.M24 + m23*self.M34)
	m34 := -(m31*self.M14 + m32*self.M24 + m33*self.M34)

	return Matrix4x4[T]{
		m11, m12, m13, m14,
		m21, m22, m23, m24,
		m31, m32, m33, m34,
		0, 0, 0, 1,
	}
}
