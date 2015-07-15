#version 330 core

#define TEXTURE_SIZE 2048

layout (points) in;
layout (triangle_strip, max_vertices = 4) out;

in mat3x2 transform[];
in vec4 texture_coordinates[];
in float texture_level[];

out vec2 texture_pos;
out float layer;

void main()
{
	layer = texture_level[0];
	vec4 texco = texture_coordinates[0];

	float width = texco.z - texco.x;
	float height = texco.w - texco.y;

	texco /= TEXTURE_SIZE;

	mat3x2 transform = transform[0];

	gl_Position.zw = vec2(0, 1);

	gl_Position.xy = transform * vec3(0, 0, 1);
	texture_pos = vec2(texco.x, texco.y);
	EmitVertex();

	gl_Position.xy = transform * vec3(0, height, 1);
	texture_pos = vec2(texco.x, texco.w);
	EmitVertex();

	gl_Position.xy = transform * vec3(width, 0, 1);
	texture_pos = vec2(texco.z, texco.y);
	EmitVertex();

	gl_Position.xy = transform * vec3(width, height, 1);
	texture_pos = vec2(texco.z, texco.w);
	EmitVertex();

	EndPrimitive();
}