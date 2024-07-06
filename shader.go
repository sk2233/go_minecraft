/*
@author: sk
@date: 2024/7/4
*/
package main

import (
	_ "embed"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sk2233/glhf"
)

var (
	//go:embed res/shader/quad.vert
	quadVertex string
	//go:embed res/shader/quad.frag
	quadFragment string
	//go:embed res/shader/chunk.vert
	chunkVertex string
	//go:embed res/shader/chunk.frag
	chunkFragment string
	//go:embed res/shader/focus.vert
	focusVertex string
	//go:embed res/shader/focus.frag
	focusFragment string
)

var (
	QuadShader  *glhf.Shader
	ChunkShader *glhf.Shader
	FocusShader *glhf.Shader
)

func InitShader() {
	var err error
	QuadShader, err = glhf.NewShader([]glhf.Attr{
		{Name: "iPos", Type: glhf.Vec3},
	}, []glhf.Attr{
		{Name: "uModel", Type: glhf.Mat4},
		{Name: "uView", Type: glhf.Mat4},
		{Name: "uProjection", Type: glhf.Mat4},
	}, quadVertex, quadFragment)
	HandleErr(err)
	QuadShader.Begin()
	QuadShader.SetUniformAttr(0, mgl32.Ident4())
	QuadShader.SetUniformAttr(2, Projection)
	QuadShader.End()

	ChunkShader, err = glhf.NewShader([]glhf.Attr{
		{Name: "iPos", Type: glhf.Vec3},
		{Name: "iTex", Type: glhf.Vec2},
		{Name: "iAo", Type: glhf.Float},
	}, []glhf.Attr{
		{Name: "uModel", Type: glhf.Mat4},
		{Name: "uView", Type: glhf.Mat4},
		{Name: "uProjection", Type: glhf.Mat4},
		{Name: "uCamera", Type: glhf.Vec3},
		{Name: "uTexture", Type: glhf.Int},
		{Name: "uSkyClr", Type: glhf.Vec3},
	}, chunkVertex, chunkFragment)
	HandleErr(err)
	ChunkShader.Begin()
	ChunkShader.SetUniformAttr(2, Projection)
	ChunkShader.SetUniformAttr(4, int32(0))
	ChunkShader.SetUniformAttr(5, mgl32.Vec3{168.0 / 255, 214.0 / 255, 250.0 / 255})
	ChunkShader.End()

	FocusShader, err = glhf.NewShader([]glhf.Attr{
		{Name: "iPos", Type: glhf.Vec3},
		{Name: "iTex", Type: glhf.Vec2},
	}, []glhf.Attr{
		{Name: "uModel", Type: glhf.Mat4},
		{Name: "uView", Type: glhf.Mat4},
		{Name: "uProjection", Type: glhf.Mat4},
		{Name: "uTexture", Type: glhf.Int},
		{Name: "uTint", Type: glhf.Vec3},
	}, focusVertex, focusFragment)
	HandleErr(err)
	FocusShader.Begin()
	FocusShader.SetUniformAttr(2, Projection)
	FocusShader.SetUniformAttr(3, int32(0))
	FocusShader.End()
}
