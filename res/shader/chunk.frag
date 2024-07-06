#version 330 core

in vec2 tex;
in float ao;
in float fogRate;

out vec4 oClr;

uniform sampler2D uTexture;
uniform vec3 uSkyClr;

void main() {
    vec4 clr = texture(uTexture,tex);
    if (clr.a==0){
        discard;
    }
    vec3 temp = mix(clr.rgb*ao, uSkyClr, fogRate);
    oClr = vec4(temp,1.0);
}