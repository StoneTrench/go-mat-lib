package path

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/StoneTrench/go-mat-lib/vec"
)

func TestPathfinder(t *testing.T) {
	type Goal struct {
		Width  int
		World  []uint8
		Target Vector3I
	}
	GetHeight := func(goal Goal) int {
		worldLen := int(len(goal.World))
		return worldLen / goal.Width
	}
	Get := func(goal Goal, x, y int) (uint8, error) {
		height := GetHeight(goal)

		if x < 0 || y < 0 || x >= goal.Width || y >= height {
			return 0, errors.New("outside the bounds of the world")
		}
		i := x + (y * goal.Width)
		return goal.World[i], nil
	}
	CanPass := func(goal Goal, x, y int) bool {
		e, err := Get(goal, x, y)
		if err != nil {
			return false
		}
		return e == 0
	}

	// 1. Setup Data
	var i uint8 = 0
	var W uint8 = 1
	world := []uint8{
		i, W, i, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		W, i, i, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		i, i, W, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, W, i, i, i, i, i, i, i, i, i, i, i,
		i, W, W, i, W, W, W, W, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		i, W, i, i, W, i, i, i, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		i, i, W, W, W, i, W, W, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		i, W, W, W, W, i, i, i, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		W, W, i, i, i, i, i, i, W, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
		W, i, i, i, i, i, i, i, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, i,
		i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i, i,
	}

	goal := Goal{
		Width:  38,
		World:  world,
		Target: Vector3I{X: 0, Y: 0, Z: 0},
	}

	// 2. Test GetHeight and Get (Bounds Checking for 100% Coverage)
	t.Run("WorldBounds", func(t *testing.T) {
		h := GetHeight(goal)
		if h != 10 {
			t.Errorf("Expected height 10, got %d", h)
		}

		// Test valid Get
		val, err := Get(goal, 0, 0)
		if err != nil || val != 0 {
			t.Error("Valid Get failed")
		}

		// Test Out of Bounds (Low)
		_, err = Get(goal, -1, 0)
		if err == nil {
			t.Error("Should have failed on X < 0")
		}

		// Test Out of Bounds (High)
		_, err = Get(goal, 38, 10)
		if err == nil {
			t.Error("Should have failed on X >= Width")
		}
	})

	// 3. Test CanPass
	t.Run("CanPass", func(t *testing.T) {
		if !CanPass(goal, 0, 0) {
			t.Error("Should be able to pass at 0,0")
		}
		if CanPass(goal, 1, 0) {
			t.Error("Should NOT be able to pass through wall (W)")
		}
		if CanPass(goal, 99, 99) {
			t.Error("Should NOT be able to pass out of bounds")
		}
	})

	// 4. Test Pathfinder Logic
	t.Run("SolvePath", func(t *testing.T) {
		path, err := GenericPathfinder(
			Vector3I{X: 0, Y: 9, Z: 0},
			goal,
			1e+9,
			func(g Goal, s Vector3I) string {
				return fmt.Sprintf("%d|%d", s.X, s.Y)
			},
			func(g Goal, s Vector3I) bool {
				return s.Equals(g.Target)
			},
			func(g Goal, s Vector3I) []Vector3I {
				neighbors := []Vector3I{
					{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1},
					{X: 1, Y: 1}, {X: 1, Y: -1}, {X: -1, Y: 1}, {X: -1, Y: -1},
				}
				var res []Vector3I
				for _, v := range neighbors {
					e := v.Add(s)
					if CanPass(g, int(e.X), int(e.Y)) {
						res = append(res, Vector3I(e))
					}
				}
				return res
			},
			func(g Goal, s, n Vector3I) float32 {
				return s.ToFloat32().Distance(n.ToFloat32())
			},
			func(g Goal, s Vector3I) float32 {
				return s.ToFloat32().Manhattan(g.Target.ToFloat32())
			},
		)

		if err != nil {
			t.Fatalf("Pathfinder returned error: %v", err)
		}

		if len(path) == 0 {
			t.Fatal("Pathfinder returned empty path")
		}

		// Verify Start and End
		if !path[0].Equals(Vector3I{X: 0, Y: 9}) || !path[len(path)-1].Equals(goal.Target) {
			t.Error("Path start or end point is incorrect")
		}
	})
}
