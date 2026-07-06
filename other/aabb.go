package other

import (
	"iter"
	"math"

	. "github.com/StoneTrench/go-mat-lib/vec"
)

type AABB[T Number] struct {
	is_not_first bool
	min          Vector3[T]
	max          Vector3[T]
}

func (self AABB[T]) Init() AABB[T] {
	self.is_not_first = false
	return self
}

func (self AABB[T]) Set(a, b Vector3[T]) AABB[T] {
	self.is_not_first = true
	self.max = a.Max(b)
	self.min = a.Min(b)
	return self
}

func (self AABB[T]) SetXYZ(x1, y1, z1, x2, y2, z2 T) AABB[T] {
	return self.Set(
		Vector3[T]{X: x1, Y: y1, Z: z1},
		Vector3[T]{X: x2, Y: y2, Z: z2},
	)
}

func (self AABB[T]) Envelop(pos Vector3[T]) AABB[T] {
	if self.is_not_first {
		self.max = self.max.Max(pos)
		self.min = self.min.Min(pos)
	} else {
		self.is_not_first = true
		self.max = pos
		self.min = pos
	}
	return self
}

func (self AABB[T]) EnvelopXYZ(x, y, z T) AABB[T] {
	return self.Envelop(Vector3[T]{X: x, Y: y, Z: z})
}

func (self AABB[T]) EnvelopAABB(other AABB[T]) AABB[T] {
	return self.Envelop(other.GetMin()).Envelop(other.GetMax())
}

func (self AABB[T]) IsInside(pos Vector3[T]) bool {
	if !self.is_not_first {
		return false
	}
	return self.min.X <= pos.X &&
		self.min.Y <= pos.Y &&
		self.min.Z <= pos.Z &&
		self.max.X >= pos.X &&
		self.max.Y >= pos.Y &&
		self.max.Z >= pos.Z
}

func (self AABB[T]) IsUnit() bool {
	size := self.GetSize()
	return size.X == 1 && size.Y == 1 && size.Z == 1
}

func (self AABB[T]) IsZero() bool {
	size := self.GetSize()
	return size.X == 0 || size.Y == 0 || size.Z == 0
}

func (self AABB[T]) GetMin() Vector3[T] {
	if !self.is_not_first {
		return Vector3[T]{X: 0, Y: 0, Z: 0}
	}
	return self.min
}

func (self AABB[T]) GetMax() Vector3[T] {
	if !self.is_not_first {
		return Vector3[T]{X: 0, Y: 0, Z: 0}
	}
	return self.max
}

func (self AABB[T]) GetSize() Vector3[T] {
	if !self.is_not_first {
		return Vector3[T]{X: 0, Y: 0, Z: 0}
	}
	return self.max.Sub(self.min).Offset(1, 1, 1)
}

func (self AABB[T]) GetCenter() Vector3[T] {
	if !self.is_not_first {
		return Vector3[T]{X: 0, Y: 0, Z: 0}
	}
	return self.max.Add(self.min).Divscale(2)
}

func (self AABB[T]) Iterate() iter.Seq[Vector3[T]] {
	return func(yield func(Vector3[T]) bool) {
		for y := self.min.Y; y <= self.max.Y; y++ {
			for x := self.min.X; x <= self.max.X; x++ {
				for z := self.min.Z; z <= self.max.Z; z++ {
					if !yield(Vector3[T]{X: x, Y: y, Z: z}) {
						return
					}
				}
			}
		}
	}
}

func (self AABB[T]) Corners() [8]Vector3[T] {
	max := self.max
	min := self.min

	return [8]Vector3[T]{
		min,
		{X: max.X, Y: min.Y, Z: min.Z},
		{X: min.X, Y: max.Y, Z: min.Z},
		{X: max.X, Y: max.Y, Z: min.Z},
		{X: min.X, Y: min.Y, Z: max.Z},
		{X: max.X, Y: min.Y, Z: max.Z},
		{X: min.X, Y: max.Y, Z: max.Z},
		max,
	}
}

func (self AABB[T]) SplitOctree() [8]AABB[T] {
	if self.IsUnit() {
		return [8]AABB[T]{
			self,
			self.Init(),
			self.Init(),
			self.Init(),
			self.Init(),
			self.Init(),
			self.Init(),
			self.Init(),
		}
	}

	corners := self.Corners()
	mid := corners[0].Add(corners[7].Sub(corners[0]).Divscale(2))
	size := self.GetSize()

	result := [8]AABB[T]{
		AABB[T]{}.Set(mid.Offset(0, 0, 0), corners[0]),
		AABB[T]{}.Set(mid.Offset(1, 0, 0), corners[1]),
		AABB[T]{}.Set(mid.Offset(0, 1, 0), corners[2]),
		AABB[T]{}.Set(mid.Offset(1, 1, 0), corners[3]),
		AABB[T]{}.Set(mid.Offset(0, 0, 1), corners[4]),
		AABB[T]{}.Set(mid.Offset(1, 0, 1), corners[5]),
		AABB[T]{}.Set(mid.Offset(0, 1, 1), corners[6]),
		AABB[T]{}.Set(mid.Offset(1, 1, 1), corners[7]),
	}

	if size.X < 2 {
		result[1] = result[1].Init()
		result[3] = result[3].Init()
		result[5] = result[5].Init()
		result[7] = result[7].Init()
	}
	if size.Y < 2 {
		result[2] = result[2].Init()
		result[3] = result[3].Init()
		result[6] = result[6].Init()
		result[7] = result[7].Init()
	}
	if size.Z < 2 {
		result[4] = result[4].Init()
		result[5] = result[5].Init()
		result[6] = result[6].Init()
		result[7] = result[7].Init()
	}

	return result
}

func (self AABB[T]) RayIntersection(entry, direction Vector3[T]) (exit Vector3[T]) {
	smin := self.GetMin()
	smax := self.GetMax()

	tmin := T(math.Inf(-1))
	tmax := T(math.Inf(1))

	if direction.X != 0 {
		t1 := (smin.X - entry.X) / direction.X
		t2 := (smax.X - entry.X) / direction.X
		tmin = max(tmin, min(t1, t2))
		tmax = min(tmax, max(t1, t2))
	} else if entry.X < smin.X || entry.X > smax.X {
		return Vector3[T]{}
	}

	if direction.Y != 0 {
		t1 := (smin.Y - entry.Y) / direction.Y
		t2 := (smax.Y - entry.Y) / direction.Y
		tmin = max(tmin, min(t1, t2))
		tmax = min(tmax, max(t1, t2))
	} else if entry.Y < smin.Y || entry.Y > smax.Y {
		return Vector3[T]{}
	}

	if direction.Z != 0 {
		t1 := (smin.Z - entry.Z) / direction.Z
		t2 := (smax.Z - entry.Z) / direction.Z
		tmin = max(tmin, min(t1, t2))
		tmax = min(tmax, max(t1, t2))
	} else if entry.Z < smin.Z || entry.Z > smax.Z {
		return Vector3[T]{}
	}

	if tmax < 0 || tmin > tmax {
		return Vector3[T]{}
	}

	tHit := tmin
	if tmin < 0 {
		tHit = 0
	}

	return entry.Add(direction.Scale(tHit))
}
