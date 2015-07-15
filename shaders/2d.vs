#version 330 core

layout(location = 0) in mat2x3 transform_in;
layout(location = 2) in vec4 texture_coordinates_in;
layout(location = 3) in float texture_level_in;

out mat3x2 transform;
out vec4 texture_coordinates;
out float texture_level;

uniform mat2x3 camera;

void main()
{
	transform = transpose(camera) * mat3(transpose(transform_in));
	texture_coordinates = texture_coordinates_in;
	texture_level = texture_level_in;
}