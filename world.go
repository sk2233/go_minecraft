/*
@author: sk
@date: 2024/7/5
*/
package main

import (
	"math/rand"
	"time"

	"github.com/sk2233/glhf"
)

type World0 struct {
	Chunks [WorldSize][WorldHeight][WorldSize]*Chunk // x y z
}

func (w *World0) Draw(shader *glhf.Shader) {
	for x := 0; x < WorldSize; x++ {
		for y := 0; y < WorldHeight; y++ {
			for z := 0; z < WorldSize; z++ {
				chunk := w.Chunks[x][y][z]
				if chunk.NeedDraw() {
					w.Chunks[x][y][z].Draw(shader)
				}
			}
		}
	}
}

func (w *World0) IsEmpty(x int, y int, z int) bool {
	// 判断世界坐标下某个位置是否为空
	if x < 0 || x >= ChunkSize*WorldSize || y < 0 || y >= ChunkSize*WorldHeight || z < 0 || z >= ChunkSize*WorldSize {
		return true
	}
	chunk := w.Chunks[x/ChunkSize][y/ChunkSize][z/ChunkSize]
	//voxel := chunk.Voxels[x%ChunkSize][y%ChunkSize][z%ChunkSize]
	return chunk.Voxels[x%ChunkSize][y%ChunkSize][z%ChunkSize] <= 0
}

func (w *World0) SetNoRefresh(x int, y int, z int, type0 VoxelType) {
	if x < 0 || x >= ChunkSize*WorldSize || y < 0 || y >= ChunkSize*WorldHeight || z < 0 || z >= ChunkSize*WorldSize {
		return
	}
	x0, y0, z0 := x/ChunkSize, y/ChunkSize, z/ChunkSize
	x1, y1, z1 := x%ChunkSize, y%ChunkSize, z%ChunkSize
	chunk := w.Chunks[x0][y0][z0]
	chunk.Voxels[x1][y1][z1] = type0
}

func (w *World0) Set(x int, y int, z int, type0 VoxelType) {
	if x < 0 || x >= ChunkSize*WorldSize || y < 0 || y >= ChunkSize*WorldHeight || z < 0 || z >= ChunkSize*WorldSize {
		return
	}
	x0, y0, z0 := x/ChunkSize, y/ChunkSize, z/ChunkSize
	x1, y1, z1 := x%ChunkSize, y%ChunkSize, z%ChunkSize
	chunk := w.Chunks[x0][y0][z0]
	chunk.Voxels[x1][y1][z1] = type0
	chunk.Refresh()
	// 关联刷新
	if x1 == 0 && x0 > 0 {
		w.Chunks[x0-1][y0][z0].Refresh()
	}
	if x1 == ChunkSize-1 && x0 < WorldSize-1 {
		w.Chunks[x0+1][y0][z0].Refresh()
	}
	if y1 == 0 && y0 > 0 {
		w.Chunks[x0][y0-1][z0].Refresh()
	}
	if y1 == ChunkSize-1 && y0 < WorldSize-1 {
		w.Chunks[x0][y0+1][z0].Refresh()
	}
	if z1 == 0 && z0 > 0 {
		w.Chunks[x0][y0][z0-1].Refresh()
	}
	if z1 == ChunkSize-1 && z0 < WorldSize-1 {
		w.Chunks[x0][y0][z0+1].Refresh()
	}
}

func (w *World0) MakeTree() {
	for i := 0; i < TreeNum; i++ {
		x := rand.Intn(WorldSize * ChunkSize)
		z := rand.Intn(WorldSize * ChunkSize)
		w.SetTree(x, z)
	}
}

func (w *World0) SetTree(x int, z int) {
	h := w.GetHeight(x, z)
	// 组装树干
	size := rand.Intn(3) + 3
	for i := 0; i < size; i++ {
		w.SetNoRefresh(x, h+i, z, VoxelWood)
	}
	h += size
	// 组装树叶
	size = min(rand.Intn(3)+2, size)
	for size >= 0 {
		for i := -size; i <= size; i++ {
			for j := -size; j <= size; j++ {
				if rand.Intn(4) > 0 {
					w.SetNoRefresh(x+i, h, z+j, VoxelLeaf)
				}
			}
		}
		size--
		h++
	}
}

func (w *World0) GetHeight(x int, z int) int {
	l, r := 0, WorldHeight*ChunkSize
	for l < r {
		m := (l + r) / 2
		if w.IsEmpty(x, m, z) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l // 这里肯定有解的
}

func (w *World0) MakeWater() {
	for x := 0; x < WorldSize*ChunkSize; x++ {
		for z := 0; z < WorldSize*ChunkSize; z++ {
			if w.IsEmpty(x, ChunkSize*3/2, z) {
				w.SetNoRefresh(x, ChunkSize*3/2, z, VoxelGlass)
			}
		}
	}
}

func NewWorld() *World0 {
	rand.Seed(time.Now().Unix())
	world0 := &World0{Chunks: [WorldSize][WorldHeight][WorldSize]*Chunk{}}
	for x := 0; x < WorldSize; x++ {
		for y := 0; y < WorldHeight; y++ {
			for z := 0; z < WorldSize; z++ {
				world0.Chunks[x][y][z] = NewChunk(world0, x*ChunkSize, y*ChunkSize, z*ChunkSize)
			}
		}
	}
	world0.MakeTree()
	//world0.MakeWater()
	return world0
}
