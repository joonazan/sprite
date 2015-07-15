#version 430 core

layout(points) in;
layout(triangle_strip, max_vertices = 4) out;

out vec2 texture_pos;
out float layer;

void main()
{
    layer = 0;

    gl_Position = vec4( 1.0, 1.0, 0.0, 1.0 );
    texture_pos = vec2( 1.0, 1.0 );
    EmitVertex();

    gl_Position = vec4(-1.0, 1.0, 0.0, 1.0 );
    texture_pos = vec2( 0.0, 1.0 ); 
    EmitVertex();

    gl_Position = vec4( 1.0,-1.0, 0.0, 1.0 );
    texture_pos = vec2( 1.0, 0.0 ); 
    EmitVertex();

    gl_Position = vec4(-1.0,-1.0, 0.0, 1.0 );
    texture_pos = vec2( 0.0, 0.0 ); 
    EmitVertex();

    EndPrimitive(); 
}