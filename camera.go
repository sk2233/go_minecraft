/*
@author: sk
@date: 2024/6/26
*/
package main

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	Camera *Camera0
)

type Camera0 struct {
	dirX, dirY float64
	pos, dir   mgl32.Vec3
	cache      mgl32.Mat4
	dirty      bool
}

func InitCamera() {
	Camera = &Camera0{dirX: -0.045235370328953195, dirY: -4.079999963790178, pos: mgl32.Vec3{486.04422, 76.66248, 12.895054}, dir: mgl32.Vec3{-0.5904688, -0.045219947, 0.8057926}, dirty: true}
}

func (c *Camera0) GetView() mgl32.Mat4 {
	if c.dirty {
		c.cache = mgl32.LookAtV(c.pos, c.pos.Add(c.dir), VecUp)
		c.dirty = false
		fmt.Println(c.dirX, c.dirY, c.pos, c.dir)
	}
	return c.cache
}

func (c *Camera0) RayCast() (xn, yn, zn, xe, ye, ze int, ok bool) { // 第一个是非空的，第二个是为空的
	pos := mgl32.Vec3{c.pos[0], c.pos[1], c.pos[2]}                                           // 原点位置
	rate := mgl32.Vec3{c.dir[0], c.dir[1], c.dir[2]}                                          // dir的比例
	dir := mgl32.Vec3{Sign(rate[0]), Sign(rate[1]), Sign(rate[2])}                            // 方向
	off := mgl32.Vec3{getOff(pos[0], dir[0]), getOff(pos[1], dir[1]), getOff(pos[2], dir[2])} // 初始偏移
	for {
		len0 := mgl32.Vec3{off[0] / rate[0], off[1] / rate[1], off[2] / rate[2]}
		if len0[0] > RayLen && len0[1] > RayLen && len0[2] > RayLen {
			return 0, 0, 0, 0, 0, 0, false // 没有找到
		}
		// 选最短的
		minIdx := 0
		if len0[1] < len0[minIdx] {
			minIdx = 1
		}
		if len0[2] < len0[minIdx] {
			minIdx = 2
		}
		x, y, z := pos[0]+len0[minIdx]*rate[0], pos[1]+len0[minIdx]*rate[1], pos[2]+len0[minIdx]*rate[2]
		xi, yi, zi := int(x), int(y), int(z)
		if minIdx == 0 {
			if !World.IsEmpty(xi, yi, zi) {
				return xi, yi, zi, xi - 1, yi, zi, true
			}
			if !World.IsEmpty(xi-1, yi, zi) {
				return xi - 1, yi, zi, xi, yi, zi, true
			}
		} else if minIdx == 1 {
			if !World.IsEmpty(xi, yi, zi) {
				return xi, yi, zi, xi, yi - 1, zi, true
			}
			if !World.IsEmpty(xi, yi-1, zi) {
				return xi, yi - 1, zi, xi, yi, zi, true
			}
		} else {
			if !World.IsEmpty(xi, yi, zi) {
				return xi, yi, zi, xi, yi, zi - 1, true
			}
			if !World.IsEmpty(xi, yi, zi-1) {
				return xi, yi, zi - 1, xi, yi, zi, true
			}
		}
		off[minIdx] += dir[minIdx]
	}
}

func getOff(base float32, dir float32) float32 {
	if dir == 0 { // 如果改方向的方向分量为 0 就放弃后续跟进
		return RayLen * 2 // 保证其长度大于 运行的长度
	}
	if base == 0 { // 当前就是一个可以检查的位置
		return 0
	}
	if dir > 0 {
		if base > 0 {
			return float32(int(base)) + 1 - base
		} else {
			return float32(int(base)) - base
		}
	} else {
		if base > 0 {
			return float32(int(base)) - base
		} else {
			return float32(int(base)) - 1 - base
		}
	}
}

func (c *Camera0) TranslateX(value float32) {
	c.pos = c.pos.Add(c.dir.Cross(VecUp).Normalize().Mul(value))
	c.dirty = true
}

func (c *Camera0) TranslateY(value float32) {
	c.pos = c.pos.Add(c.dir.Cross(VecRight).Normalize().Mul(value))
	c.dirty = true
}

func (c *Camera0) TranslateZ(value float32) {
	c.pos = c.pos.Add(c.dir.Normalize().Mul(value))
	c.dirty = true
}

func (c *Camera0) RotateX(value float32) { // 左右看 沿着 Y轴
	c.dirY -= float64(value)
	c.updateDir()
}

func (c *Camera0) RotateY(value float32) { // 上下看 沿着  X轴
	c.dirX += float64(value)
	c.updateDir()
}

func (c *Camera0) updateDir() {
	c.dir = mgl32.Vec3{
		float32(math.Cos(c.dirX) * math.Cos(c.dirY)),
		float32(math.Sin(c.dirX)),
		float32(math.Cos(c.dirX) * math.Sin(c.dirY)),
	}
	c.dirty = true
}

func (c *Camera0) GetPos() mgl32.Vec3 {
	return c.pos
}

func (c *Camera0) InView(pos mgl32.Vec3, r float32) bool {
	pos = pos.Sub(c.pos)
	l := pos.Dot(c.dir) // dir 的长度本来就为 1 不用除了
	if l-r > ViewFar || l+r < ViewNear {
		return false
	}
	return true
}
