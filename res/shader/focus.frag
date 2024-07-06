#version 330 core

in vec2 tex;

out vec4 oClr;

uniform sampler2D uTexture;
uniform vec3 uTint;

void main() {
    vec4 clr = texture(uTexture,tex);
    if (clr.a==0){
        discard;
    }
    oClr = vec4(clr.rgb*uTint,1.0);
}