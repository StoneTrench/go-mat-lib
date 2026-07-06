package grid

import (
	"iter"

	. "github.com/StoneTrench/go-mat-lib/other"
	. "github.com/StoneTrench/go-mat-lib/vec"
)

const CHUNK_SIZE_EXP int32 = 5
const CHUNK_SIZE int32 = 1 << CHUNK_SIZE_EXP
const CHUNK_SIZE_MASK = (CHUNK_SIZE - 1)

type Chunk[T any] struct {
	Cells [CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]T
}
type Grid[T any] struct {
	AABB   AABB[int32]
	Chunks map[Vector3I]*Chunk[T]
}

func (grid Grid[T]) Init() *Grid[T] {
	grid.AABB = grid.AABB.Init()
	grid.Chunks = make(map[Vector3I]*Chunk[T])

	return &grid
}

func (grid *Grid[T]) GetPositionInfo(pos Vector3I) (chunk_pos Vector3I, local_pos Vector3I, chunk *Chunk[T]) {
	chunk_pos = pos.RShift32(CHUNK_SIZE_EXP)
	lx := pos.X & CHUNK_SIZE_MASK
	ly := pos.Y & CHUNK_SIZE_MASK
	lz := pos.Z & CHUNK_SIZE_MASK
	local_pos = Vector3I{X: lx, Y: ly, Z: lz}

	if grid.Chunks == nil {
		grid.Chunks = map[Vector3I]*Chunk[T]{}
	}

	chunk = grid.Chunks[chunk_pos]
	if chunk == nil {
		chunk = &Chunk[T]{
			[CHUNK_SIZE][CHUNK_SIZE][CHUNK_SIZE]T{},
		}
		grid.Chunks[chunk_pos] = chunk
	}

	return chunk_pos, local_pos, chunk
}

func (grid *Grid[T]) GetCell(pos Vector3I) T {
	_, local_pos, chunk := grid.GetPositionInfo(pos)
	return chunk.Cells[local_pos.X][local_pos.Y][local_pos.Z]
}

func (grid *Grid[T]) SetCell(pos Vector3I, cell T) {
	_, local_pos, chunk := grid.GetPositionInfo(pos)
	chunk.Cells[local_pos.X][local_pos.Y][local_pos.Z] = cell
	grid.AABB = grid.AABB.Envelop(pos)
}

func (grid *Grid[T]) SetCellXYZ(x, y, z int32, cell T) {
	grid.SetCell(Vector3I{X: x, Y: y, Z: z}, cell)
}

func (grid *Grid[T]) Iterate() iter.Seq2[Vector3I, T] {
	return func(yield func(Vector3I, T) bool) {
		for chunk_pos, chunk := range grid.Chunks {
			for k, v := range grid.IterateChunkRaw(chunk_pos, chunk) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func (grid *Grid[T]) IterateChunkRaw(chunk_pos Vector3I, chunk *Chunk[T]) iter.Seq2[Vector3I, T] {
	return func(yield func(Vector3I, T) bool) {
		world_pos := chunk_pos.LShift32(CHUNK_SIZE_EXP)
		for x, bx := range chunk.Cells {
			for y, by := range bx {
				for z, Id := range by {
					pos := world_pos.Offset(int32(x), int32(y), int32(z))
					if !yield(pos, Id) {
						return
					}
				}
			}
		}
	}
}

func (grid *Grid[T]) IterateChunk(chunk_pos Vector3I) iter.Seq2[Vector3I, T] {
	return grid.IterateChunkRaw(chunk_pos, grid.Chunks[chunk_pos])
}
