#version 330 core

in vec3 iPos;
in vec2 iTex;

out vec2 tex;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main() {
    tex=iTex;
    gl_Position = uProjection * uView * uModel * vec4(iPos, 1.0);
}