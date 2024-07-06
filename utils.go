/*
@author: sk
@date: 2024/7/3
*/
package main

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"time"

	"github.com/aquilax/go-perlin"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sk2233/glhf"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PressKey(window *glfw.Window, key glfw.Key) bool {
	return window.GetKey(key) == glfw.Press // 查询的是状态值
}

func ReleaseKey(window *glfw.Window, key glfw.Key) bool {
	return window.GetKey(key) == glfw.Release // 查询的是状态值
}

func GetAxis(window *glfw.Window, min, max glfw.Key) float32 {
	if PressKey(window, min) {
		return -1
	}
	if PressKey(window, max) {
		return 1
	}
	return 0
}

func Append(data []float32, items ...[]float32) []float32 {
	for _, item := range items {
		data = append(data, item...)
	}
	return data
}

func LoadTexture(path string) *glhf.Texture {
	file, err := os.Open(path)
	HandleErr(err)
	img, _, err := image.Decode(file)
	HandleErr(err)

	bound := img.Bounds()
	temp := image.NewNRGBA(bound)
	draw.Draw(temp, temp.Bounds(), img, bound.Min, draw.Src)
	return glhf.NewTexture(bound.Dx(), bound.Dy(), false, temp.Pix)
}

func GetOffset(voxel VoxelType, face FaceType) (float32, float32) {
	switch voxel {
	case VoxelGrass:
		switch face {
		case FaceTop:
			return 0, 13 * VoxelSize
		case FaceBottom:
			return 0, 15 * VoxelSize
		default:
			return 0, 14 * VoxelSize
		}
	case VoxelSand:
		return VoxelSize, 15 * VoxelSize
	case VoxelStone:
		return 12 * VoxelSize, 15 * VoxelSize
	case VoxelBrick:
		return 3 * VoxelSize, 15 * VoxelSize
	case VoxelWood:
		switch face {
		case FaceTop:
			return 4 * VoxelSize, 13 * VoxelSize
		case FaceBottom:
			return 4 * VoxelSize, 15 * VoxelSize
		default:
			return 4 * VoxelSize, 14 * VoxelSize
		}
	case VoxelSoil:
		return 6 * VoxelSize, 15 * VoxelSize
	case VoxelGlass:
		return 9 * VoxelSize, 15 * VoxelSize
	case VoxelLeaf:
		return 14 * VoxelSize, 15 * VoxelSize
	case VoxelSnow:
		switch face {
		case FaceTop:
			return 8 * VoxelSize, 13 * VoxelSize
		case FaceBottom:
			return 8 * VoxelSize, 15 * VoxelSize
		default:
			return 8 * VoxelSize, 14 * VoxelSize
		}
	default:
		return 15 * VoxelSize, 15 * VoxelSize
	}
}

var (
	perlin0 = perlin.NewPerlin(2, 2, 4, time.Now().Unix())
)

func Noise(x, y float32) float32 {
	return float32(perlin0.Noise2D(float64(x), float64(y)))
}

func Sign(val float32) float32 {
	if val < 0 {
		return -1
	} else if val > 0 {
		return 1
	} else {
		return 0
	}
}

var (
	randTypes = []VoxelType{
		VoxelSand,
		VoxelSoil, VoxelSoil,
		VoxelGrass, VoxelGrass, VoxelGrass, VoxelGrass,
		VoxelSnow, VoxelSnow, VoxelSnow, VoxelSnow, VoxelSnow, VoxelSnow, VoxelSnow, VoxelSnow,
	}
)

func GetRandVoxel(h int, max int) VoxelType {
	if h > WorldHeight*ChunkSize*2/3 {
		if h == max-1 {
			return VoxelSnow
		} else {
			return VoxelSoil
		}
		//return randTypes[rand.Intn(1+2+4+8)]
	} else if h > WorldHeight*ChunkSize/3 {
		if h == max-1 {
			return VoxelGrass
		} else {
			return VoxelSoil
		}
		//return randTypes[rand.Intn(1+2+4)]
	} else {
		return VoxelSand
	}
}
