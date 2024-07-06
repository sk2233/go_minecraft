/*
@author: sk
@date: 2024/7/3
*/
package main

import "github.com/go-gl/mathgl/mgl32"

const (
	WinW = 1280.0
	WinH = 720.0
)

var (
	VecFront = mgl32.Vec3{0, 0, -1}
	VecUp    = mgl32.Vec3{0, 1, 0}
	VecRight = mgl32.Vec3{1, 0, 0}
)

const (
	ChunkSize   = 32
	WorldSize   = 16
	WorldHeight = 4
)

type VoxelType int

const (
	VoxelNone  VoxelType = 0
	VoxelGrass VoxelType = 1
	VoxelSand  VoxelType = 2
	VoxelStone VoxelType = 3
	VoxelBrick VoxelType = 4
	VoxelWood  VoxelType = 5
	VoxelSoil  VoxelType = 6
	VoxelGlass VoxelType = 7
	VoxelLeaf  VoxelType = 8
	VoxelSnow  VoxelType = 9
)

type FaceType int

const (
	FaceTop    FaceType = 1 // 方便放进去
	FaceBottom FaceType = 2
	FaceRight  FaceType = 3
	FaceLeft   FaceType = 4
	FaceBack   FaceType = 5
	FaceFront  FaceType = 6
)

type PlaneType int

const (
	PlaneX PlaneType = 1
	PlaneY PlaneType = 2
	PlaneZ PlaneType = 3
)

const (
	VoxelCount = 16 // 16 * 16 个纹理
	VoxelSize  = 1.0 / VoxelCount
)

const (
	RayLen = 8 // 射线长度
)

const (
	ViewNear = 0.1
	ViewFar  = 2000
)

const (
	TreeNum = 100
)
