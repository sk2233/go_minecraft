#version 330 core

in vec3 iPos;
in vec2 iTex;
in float iAo;

out vec2 tex;
out float ao;
out float fogRate;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;
uniform vec3 uCamera;

void main() {
    tex=iTex;
    ao=(iAo+1.0)*0.25;
    vec4 pos = uModel*vec4(iPos,1.0); // 必须使用世界坐标
    float len = distance(pos.xyz, uCamera);
    fogRate = pow(clamp(len/(6*32), 0, 1), 4);
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}