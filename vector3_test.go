package lib

import (
	"reflect"
	"testing"
)

func TestVector3(t *testing.T) {
	v1 := Vector3[float64]{X: 1, Y: 2, Z: 2}
	v2 := Vector3[float64]{X: 3, Y: 4, Z: 5}

	t.Run("Setters", func(t *testing.T) {
		if res := v1.Set(10, 11, 12); res.X != 10 || res.Y != 11 || res.Z != 12 {
			t.Error("Set failed")
		}
		if res := v1.SetX(9); res.X != 9 {
			t.Error("SetX failed")
		}
		if res := v1.SetY(9); res.Y != 9 {
			t.Error("SetY failed")
		}
		if res := v1.SetZ(9); res.Z != 9 {
			t.Error("SetZ failed")
		}
		if res := v1.Offset(1, 1, 1); res.X != 2 {
			t.Error("Offset failed")
		}
	})

	t.Run("Arithmetic", func(t *testing.T) {
		if res := v1.Add(v2); res != (Vector3[float64]{4, 6, 7}) {
			t.Error("Add failed")
		}
		if res := v2.Sub(v1); res != (Vector3[float64]{2, 2, 3}) {
			t.Error("Sub failed")
		}
		if res := v1.Scale(2); res != (Vector3[float64]{2, 4, 4}) {
			t.Error("Scale failed")
		}
		if res := v1.Divscale(2); res != (Vector3[float64]{0.5, 1, 1}) {
			t.Error("Divscale failed")
		}
		if res := v1.Negate(); res != (Vector3[float64]{-1, -2, -2}) {
			t.Error("Negate failed")
		}
		if res := (Vector3[float64]{-1, 2, -3}).Abs(); res != (Vector3[float64]{1, 2, 3}) {
			t.Error("Abs failed")
		}
	})

	t.Run("MathProducts", func(t *testing.T) {
		if dot := v1.Dot(v2); dot != 21 {
			t.Errorf("Dot failed: got %f", dot)
		}
		if res := v1.Cross(v2); res != (Vector3[float64]{2, 1, -2}) {
			t.Error("Cross failed")
		}
		if dist := v1.Distance(Vector3[float64]{1, 2, 5}); dist != 3 {
			t.Error("Distance failed")
		}
		if len := v1.Length(); len != 3 {
			t.Error("Length failed")
		}
		if res := v1.Normalize(); res.Length() != 1 {
			t.Error("Normalize failed")
		}
		if res := v1.Apply(func(x float64) float64 { return x * x }); res.X != 1 || res.Y != 4 {
			t.Error("Apply failed")
		}
	})

	t.Run("Conversions", func(t *testing.T) {
		slice := v1.ToSlice()
		if len(slice) != 3 || slice[0] != 1 {
			t.Error("ToSlice failed")
		}
		fixed := v1.ToFixedSlice()
		if fixed[1] != 2 {
			t.Error("ToFixedSlice failed")
		}

		v3 := Vector3[float64]{}.FromSlice([]float64{7, 8, 9})
		if v3.X != 7 {
			t.Error("FromSlice failed")
		}
		v4 := Vector3[float64]{}.FromFixedSlice([3]float64{10, 11, 12})
		if v4.Z != 12 {
			t.Error("FromFixedSlice failed")
		}
	})

	t.Run("Reductions", func(t *testing.T) {
		if v1.Sum() != 5 {
			t.Error("Sum failed")
		}
		if v1.Volume() != 4 {
			t.Error("Volume failed")
		}
		if v1.InnerProduct() != 9 {
			t.Error("InnerProduct (Self-Dot) failed")
		}
		if v1.Manhattan(v2) != 7 {
			t.Error("Manhattan failed")
		}
		if !v1.Equals(v1) {
			t.Error("Equals failed")
		}
	})

	t.Run("MinMax", func(t *testing.T) {
		min, max := v1.Minmax(v2, Vector3[float64]{0, 10, 0})
		if min.X != 0 || max.Y != 10 {
			t.Error("Minmax variadic failed")
		}
		if v1.Max(v2) != v2 {
			t.Error("Max failed")
		}
		if v1.Min(v2) != v1 {
			t.Error("Min failed")
		}
	})

	t.Run("TypeCasting", func(t *testing.T) {
		vf := Vector3[float64]{1.1, 2.5, 3.9}
		if res := vf.FloorToInt32(); res.X != 1 || res.Z != 3 {
			t.Error("Floor failed")
		}
		if res := vf.CeilToInt32(); res.X != 2 || res.Z != 4 {
			t.Error("Ceil failed")
		}
		if res := vf.RoundToInt32(); res.X != 1 || res.Y != 3 || res.Z != 4 {
			t.Error("Round failed")
		}

		if reflect.TypeOf(v1.ToFloat32().X).Kind() != reflect.Float32 {
			t.Error("ToFloat32 failed")
		}
		if reflect.TypeOf(v1.ToFloat64().X).Kind() != reflect.Float64 {
			t.Error("ToFloat64 failed")
		}
	})

	t.Run("Shifting", func(t *testing.T) {
		vi := Vector3[int32]{2, 4, 8}
		if res := vi.LShift32(1); res.X != 4 {
			t.Error("LShift32 failed")
		}
		if res := vi.RShift32(1); res.X != 1 {
			t.Error("RShift32 failed")
		}
	})

	t.Run("Matrix", func(t *testing.T) {
		var m Matrix4x4[float64]
		_ = v1.MulMat(m)
	})
}
