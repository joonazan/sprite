#version 430 core

in vec2 texture_pos;
in float layer;

layout(binding = 0) uniform sampler2DArray taustat;

layout(location = 0) out vec4 color;


void main()
{
	color = texture(taustat, vec3(texture_pos, layer));
}