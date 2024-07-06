/*
@author: sk
@date: 2024/7/5
*/
package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sk2233/glhf"
)

type Chunk struct {
	Voxels  [ChunkSize][ChunkSize][ChunkSize]VoxelType
	Mesh    *glhf.VertexSlice
	World   *World0
	X, Y, Z int
}

func NewChunk(world *World0, x, y, z int) *Chunk {
	voxels := [ChunkSize][ChunkSize][ChunkSize]VoxelType{}
	for x0 := 0; x0 < ChunkSize; x0++ {
		for z0 := 0; z0 < ChunkSize; z0++ {
			x1 := float32(x0+x) / (ChunkSize * WorldSize)
			z1 := float32(z0+z) / (ChunkSize * WorldSize)
			// 入参必须是 0 ~ 1 返回值是 -1 ～ 1
			h := ((Noise(x1, z1) + 1) / 2) * WorldHeight * ChunkSize
			count := min(ChunkSize, int(h)-y)
			for y0 := 0; y0 < count; y0++ {
				voxels[x0][y0][z0] = GetRandVoxel(y0+y, int(h))
			}
		}
	}
	return &Chunk{Voxels: voxels, World: world, X: x, Y: y, Z: z}
}

func (c *Chunk) GetMesh() *glhf.VertexSlice {
	if c.Mesh == nil {
		data := make([]float32, 0)
		for x := 0; x < ChunkSize; x++ {
			for y := 0; y < ChunkSize; y++ {
				for z := 0; z < ChunkSize; z++ {
					if c.Voxels[x][y][z] > 0 {
						type0 := c.Voxels[x][y][z]
						// 有内容开始收集，忽略隐藏面
						if c.IsEmpty(x, y+1, z) {
							x0, y0 := GetOffset(type0, FaceTop)
							a0, a1, a2, a3 := c.GetAo(x, y+1, z, PlaneY)
							v0 := []float32{float32(x), float32(y + 1), float32(z), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x + 1), float32(y + 1), float32(z), x0, y0, a1}
							v2 := []float32{float32(x + 1), float32(y + 1), float32(z + 1), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x), float32(y + 1), float32(z + 1), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v3, v2, v0, v2, v1)
						}
						if c.IsEmpty(x, y-1, z) {
							x0, y0 := GetOffset(type0, FaceBottom)
							a0, a1, a2, a3 := c.GetAo(x, y-1, z, PlaneY)
							v0 := []float32{float32(x), float32(y), float32(z), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x + 1), float32(y), float32(z), x0, y0, a1}
							v2 := []float32{float32(x + 1), float32(y), float32(z + 1), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x), float32(y), float32(z + 1), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v2, v3, v0, v1, v2)
						}
						if c.IsEmpty(x+1, y, z) {
							x0, y0 := GetOffset(type0, FaceRight)
							a0, a1, a2, a3 := c.GetAo(x+1, y, z, PlaneX)
							v0 := []float32{float32(x + 1), float32(y), float32(z), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x + 1), float32(y + 1), float32(z), x0, y0, a1}
							v2 := []float32{float32(x + 1), float32(y + 1), float32(z + 1), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x + 1), float32(y), float32(z + 1), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v1, v2, v0, v2, v3)
						}
						if c.IsEmpty(x-1, y, z) {
							x0, y0 := GetOffset(type0, FaceLeft)
							a0, a1, a2, a3 := c.GetAo(x-1, y, z, PlaneX)
							v0 := []float32{float32(x), float32(y), float32(z), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x), float32(y + 1), float32(z), x0, y0, a1}
							v2 := []float32{float32(x), float32(y + 1), float32(z + 1), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x), float32(y), float32(z + 1), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v2, v1, v0, v3, v2)
						}
						if c.IsEmpty(x, y, z-1) {
							x0, y0 := GetOffset(type0, FaceBack)
							a0, a1, a2, a3 := c.GetAo(x, y, z-1, PlaneZ)
							v0 := []float32{float32(x), float32(y), float32(z), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x), float32(y + 1), float32(z), x0, y0, a1}
							v2 := []float32{float32(x + 1), float32(y + 1), float32(z), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x + 1), float32(y), float32(z), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v1, v2, v0, v2, v3)
						}
						if c.IsEmpty(x, y, z+1) {
							x0, y0 := GetOffset(type0, FaceFront)
							a0, a1, a2, a3 := c.GetAo(x, y, z+1, PlaneZ)
							v0 := []float32{float32(x), float32(y), float32(z + 1), x0, y0 + VoxelSize, a0}
							v1 := []float32{float32(x), float32(y + 1), float32(z + 1), x0, y0, a1}
							v2 := []float32{float32(x + 1), float32(y + 1), float32(z + 1), x0 + VoxelSize, y0, a2}
							v3 := []float32{float32(x + 1), float32(y), float32(z + 1), x0 + VoxelSize, y0 + VoxelSize, a3}
							data = Append(data, v0, v2, v1, v0, v3, v2)
						}
					}
				}
			}
		}
		c.Mesh = glhf.MakeVertexSlice(ChunkShader, len(data)/6, len(data)/6)
		c.Mesh.Begin()
		c.Mesh.SetVertexData(data)
		c.Mesh.End()
	}
	return c.Mesh
}

func (c *Chunk) IsEmpty(x int, y int, z int) bool {
	return c.World.IsEmpty(c.X+x, c.Y+y, c.Z+z)
}

func (c *Chunk) Draw(shader *glhf.Shader) {
	shader.SetUniformAttr(0, mgl32.Translate3D(float32(c.X), float32(c.Y), float32(c.Z)))
	mesh := c.GetMesh()
	mesh.Begin()
	mesh.Draw()
	mesh.End()
}

// 简单计算环境遮蔽光 根据其周围的8个方块确定
func (c *Chunk) GetAo(x int, y int, z int, type0 PlaneType) (float32, float32, float32, float32) {
	a, b, c0, d, e, f, g, h := float32(0), float32(0), float32(0), float32(0), float32(0), float32(0), float32(0), float32(0)
	if type0 == PlaneX {
		if c.IsEmpty(x, y, z-1) {
			a = 1
		}
		if c.IsEmpty(x, y-1, z-1) {
			b = 1
		}
		if c.IsEmpty(x, y-1, z) {
			c0 = 1
		}
		if c.IsEmpty(x, y-1, z+1) {
			d = 1
		}
		if c.IsEmpty(x, y, z+1) {
			e = 1
		}
		if c.IsEmpty(x, y+1, z+1) {
			f = 1
		}
		if c.IsEmpty(x, y+1, z) {
			g = 1
		}
		if c.IsEmpty(x, y+1, z-1) {
			h = 1
		}
	} else if type0 == PlaneY {
		if c.IsEmpty(x, y, z-1) {
			a = 1
		}
		if c.IsEmpty(x-1, y, z-1) {
			b = 1
		}
		if c.IsEmpty(x-1, y, z) {
			c0 = 1
		}
		if c.IsEmpty(x-1, y, z+1) {
			d = 1
		}
		if c.IsEmpty(x, y, z+1) {
			e = 1
		}
		if c.IsEmpty(x+1, y, z+1) {
			f = 1
		}
		if c.IsEmpty(x+1, y, z) {
			g = 1
		}
		if c.IsEmpty(x+1, y, z-1) {
			h = 1
		}
	} else {
		if c.IsEmpty(x-1, y, z) {
			a = 1
		}
		if c.IsEmpty(x-1, y-1, z) {
			b = 1
		}
		if c.IsEmpty(x, y-1, z) {
			c0 = 1
		}
		if c.IsEmpty(x+1, y-1, z) {
			d = 1
		}
		if c.IsEmpty(x+1, y, z) {
			e = 1
		}
		if c.IsEmpty(x+1, y+1, z) {
			f = 1
		}
		if c.IsEmpty(x, y+1, z) {
			g = 1
		}
		if c.IsEmpty(x-1, y+1, z) {
			h = 1
		}
	}
	return a + b + c0, g + h + a, e + f + g, c0 + d + e
}

func (c *Chunk) Refresh() {
	c.Mesh = nil // 会自动创建
}

var (
	Sqrt3 = float32(math.Sqrt(3))
)

func (c *Chunk) NeedDraw() bool {
	return Camera.InView(mgl32.Vec3{float32(c.X + ChunkSize/2), float32(c.Y + ChunkSize/2), float32(c.Z + ChunkSize/2)}, Sqrt3*ChunkSize/2)
}
