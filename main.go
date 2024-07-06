/*
@author: sk
@date: 2024/7/3
*/
package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopxl/mainthread/v2"
	"github.com/sk2233/glhf"
)

var (
	Projection = mgl32.Perspective(mgl32.DegToRad(45.0), WinW/WinH, ViewNear, ViewFar)

	Window *glfw.Window

	World                     = NewWorld()
	AddMode                   = true // 添加方块模式，否则就是移除方块模式
	Select                    bool
	SelectX, SelectY, SelectZ int
)

func InitData() {
	// 准备基础环境
	mainthread.Call(func() {
		// gl环境设置
		err := glfw.Init()
		HandleErr(err)
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.Resizable, glfw.False)
		// 窗口设置
		Window, err = glfw.CreateWindow(WinW, WinH, "Main", nil, nil)
		HandleErr(err)
		Window.MakeContextCurrent()
		Window.SetKeyCallback(HandleKey)
		glhf.Init()
		// 全局设置
		gl.Enable(gl.CULL_FACE)  // 开启剔除背面
		gl.Enable(gl.DEPTH_TEST) // 开启深度检测
		gl.DepthFunc(gl.LESS)
		//gl.Enable(gl.BLEND) // 开启透明  直接丢弃透明颜色
		//gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.ClearColor(168.0/255, 214.0/255, 250.0/255, 1)
		// 其他设置
		InitShader()
		InitMesh()
		InitCamera()
		InitTexture()
	})
}

func MainLoop() {
	mainthread.Call(func() {
		for !Window.ShouldClose() {
			Update()
			Draw()

			Window.SwapBuffers()
			glfw.PollEvents()
		}
	})
}

func Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT) // 先应用全局设置清楚相关内容

	ChunkShader.Begin()
	ChunkShader.SetUniformAttr(1, Camera.GetView())
	ChunkShader.SetUniformAttr(3, Camera.GetPos())
	VoxelsTexture.Begin()
	World.Draw(ChunkShader)
	VoxelsTexture.End()
	ChunkShader.End()

	if Select {
		FocusShader.Begin()
		FocusShader.SetUniformAttr(1, Camera.GetView())
		FocusShader.SetUniformAttr(0, mgl32.Translate3D(float32(SelectX), float32(SelectY), float32(SelectZ)))
		if AddMode {
			FocusShader.SetUniformAttr(4, mgl32.Vec3{0, 0, 1})
		} else {
			FocusShader.SetUniformAttr(4, mgl32.Vec3{1, 0, 0})
		}
		FocusTexture.Begin()
		FocusMesh.Begin()
		FocusMesh.Draw()
		FocusMesh.End()
		FocusTexture.End()
		FocusShader.End()
	}
}

func Update() {
	HandleInput()
	HandleFocus()
}

func HandleFocus() {
	if xn, yn, zn, xe, ye, ze, ok := Camera.RayCast(); ok {
		Select = true
		if AddMode {
			SelectX, SelectY, SelectZ = xe, ye, ze
		} else {
			SelectX, SelectY, SelectZ = xn, yn, zn
		}
	} else {
		Select = false
	}
}

func HandleInput() {
	// 位移
	offsetX := GetAxis(Window, glfw.KeyA, glfw.KeyD)
	if offsetX != 0 {
		Camera.TranslateX(offsetX * 0.1)
	}
	offsetY := GetAxis(Window, glfw.KeyE, glfw.KeyQ)
	if offsetY != 0 {
		Camera.TranslateY(offsetY * 0.1)
	}
	offsetZ := GetAxis(Window, glfw.KeyS, glfw.KeyW)
	if offsetZ != 0 {
		Camera.TranslateZ(offsetZ * 0.1)
	}
	// 旋转
	rotateX := GetAxis(Window, glfw.KeyRight, glfw.KeyLeft)
	if rotateX != 0 {
		Camera.RotateX(rotateX * 0.01)
	}
	rotateY := GetAxis(Window, glfw.KeyDown, glfw.KeyUp)
	if rotateY != 0 {
		Camera.RotateY(rotateY * 0.01)
	}
}

func HandleKey(_ *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}
	switch key {
	case glfw.KeyEscape:
		Window.SetShouldClose(true)
	case glfw.KeyR: // 切换模式
		AddMode = !AddMode
	case glfw.KeySpace: // 修改聚焦模块
		if Select {
			if AddMode {
				World.Set(SelectX, SelectY, SelectZ, VoxelGlass)
			} else {
				World.Set(SelectX, SelectY, SelectZ, VoxelNone)
			}
		}
	}
}

func EndData() {
	mainthread.Call(func() {
		glfw.Terminate()
	})
}

func Main() {
	InitData()
	MainLoop()
	EndData()
}

// https://www.youtube.com/watch?v=Ab8TOSFfNp4

func main() {
	mainthread.Run(Main)
}
