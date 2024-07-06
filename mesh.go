/*
@author: sk
@date: 2024/7/4
*/
package main

import "github.com/sk2233/glhf"

var (
	QuadMesh  *glhf.VertexSlice
	FocusMesh *glhf.VertexSlice
)

func InitMesh() {
	QuadMesh = glhf.MakeVertexSlice(QuadShader, 6, 6)
	QuadMesh.Begin()
	QuadMesh.SetVertexData(QuadData)
	QuadMesh.End()

	FocusMesh = glhf.MakeVertexSlice(FocusShader, 36, 36)
	FocusMesh.Begin()
	FocusMesh.SetVertexData(FocusData)
	FocusMesh.End()
}
