/*
@author: sk
@date: 2024/7/5
*/
package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/sk2233/glhf"
)

var (
	BlockTexture  *glhf.Texture
	VoxelsTexture *glhf.Texture
	FocusTexture  *glhf.Texture
)

func InitTexture() {
	BlockTexture = LoadTexture("res/texture/block.png")
	BlockTexture.SetTextureIdx(gl.TEXTURE0)
	VoxelsTexture = LoadTexture("res/texture/voxels.png")
	VoxelsTexture.SetTextureIdx(gl.TEXTURE0)
	FocusTexture = LoadTexture("res/texture/focus.png")
	FocusTexture.SetTextureIdx(gl.TEXTURE0)
}
