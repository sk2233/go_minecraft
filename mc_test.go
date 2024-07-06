/*
@author: sk
@date: 2024/7/5
*/
package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strings"
	"testing"

	"golang.org/x/image/colornames"
)

func TestChunk(t *testing.T) {
	//chunk := NewChunk()
	//chunk.Voxels[12][12][12] = 0
}

func TestImg(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			if x == 0 || y == 0 || x == 15 || y == 15 {
				//img.Set(x, y, colornames.Black)
				img.Set(x, y, colornames.White)
			}
		}
	}
	file, err := os.Create("/Users/bytedance/Documents/go/minecraft/res/texture/focus.png")
	HandleErr(err)
	png.Encode(file, img)
}

func TestImg2(t *testing.T) {
	file, err := os.Open("/Users/bytedance/Documents/go/minecraft/res/texture/voxels.png")
	HandleErr(err)
	src, err := png.Decode(file)
	HandleErr(err)

	dst := image.NewRGBA(src.Bounds())
	for i := 0; i < src.Bounds().Dx(); i++ {
		for j := 0; j < src.Bounds().Dy(); j++ {
			clr := src.At(i, j)
			// 234, 51, 247
			r, g, b, _ := clr.RGBA()
			if r == 65535 && g == 0 && b == 65535 {
				continue
			}
			dst.Set(i, j, clr)
		}
	}
	file, err = os.Create("/Users/bytedance/Documents/go/minecraft/res/texture/test.png")
	HandleErr(err)
	png.Encode(file, dst)
}

func TestNoise(t *testing.T) {
	for i := 0; i < 1000; i++ {
		//fmt.Println(Noise(rand.Float64(), rand.Float64()))
	}
}

func TestNum(t *testing.T) {
	num := -22.33
	fmt.Println(int(num))
}

func TestMesh(t *testing.T) {
	buff := &strings.Builder{}
	for i := 0; i < len(OldFocusData); i += 5 {
		buff.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,\n", Adjust(FocusData[i]), Adjust(FocusData[i+1]), Adjust(FocusData[i+2]),
			FocusData[i+3], FocusData[i+4]))
	}
	fmt.Println(buff.String())
}

func Adjust(num float32) float32 {
	if num == 0 {
		return -0.01
	}
	if num == 1 {
		return 1.01
	}
	return 0
}

func TestNum2(t *testing.T) {
	fmt.Println(math.Sqrt(3) * math.Sqrt(3))
}
