package sprite

import (
	"github.com/go-gl-legacy/gl"
	glfw "github.com/go-gl/glfw3/v3.0/glfw"
	"github.com/joonazan/vec2"
	"unsafe"
)

type Image struct {
	TextureLeft, TextureTop     float32
	TextureRight, TextureBottom float32
	Layer                       float32
}

type Sprite struct {
	Transform [6]float32
	Image
}

type SpriteDrawer struct {
	gl.Program
	Camera         vec2.Matrix
	camera_uniform gl.UniformLocation
	Texture        gl.Texture
	Window         *glfw.Window

	width, height  int
	viewportMatrix vec2.Matrix
}

func NewSpriteDrawer(window *glfw.Window, layers int) *SpriteDrawer {
	s := new(SpriteDrawer)

	vao := gl.GenVertexArray()
	vao.Bind()

	s.Camera = vec2.Identity

	s.Program = CreateProgram("shaders/2d.vert", "shaders/2d.geom", "shaders/texture.frag")
	s.Use()
	s.camera_uniform = s.GetUniformLocation("camera")

	s.Texture = gl.GenTexture()
	s.Texture.Bind(gl.TEXTURE_2D_ARRAY)
	gl.TexImage3D(gl.TEXTURE_2D_ARRAY, 0, gl.RGBA, 2048, 2048, layers, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D_ARRAY, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	s.Window = window
	w, h := window.GetSize()
	s.OnScreenResize(w, h)

	s.Window.SetSizeCallback(func(window *glfw.Window, w, h int) {
		s.OnScreenResize(w, h)
	})

	return s
}

func (d *SpriteDrawer) OnScreenResize(width, height int) {
	d.width = width
	d.height = height
	d.viewportMatrix = vec2.Scale(2.0/float64(width), 2.0/float64(height))
}

func (d *SpriteDrawer) GetTransform() vec2.Matrix {
	return d.Camera.Mul(d.viewportMatrix)
}

func (d *SpriteDrawer) GetMousePos() vec2.Vector {
	x, y := d.Window.GetCursorPosition()

	screen := vec2.Vector{x/float64(d.width)*2 - 1, -y/float64(d.height)*2 + 1}

	inv_transform := d.GetTransform().Inverse()
	return inv_transform.Transform(screen)
}

func (drawer *SpriteDrawer) Draw(sprites []Sprite) {

	gl.Viewport(0, 0, drawer.width, drawer.height)

	if len(sprites) == 0 {
		return
	}

	drawer.Use()
	drawer.Texture.Bind(gl.TEXTURE_2D_ARRAY)

	tmp := drawer.GetTransform().To32()
	drawer.camera_uniform.UniformMatrix2x3f(false, &tmp)

	vertexbuffer := gl.GenBuffer()
	defer vertexbuffer.Delete()
	vertexbuffer.Bind(gl.ARRAY_BUFFER)

	stride := int(unsafe.Sizeof(sprites[0]))

	gl.BufferData(gl.ARRAY_BUFFER, stride*len(sprites), sprites, gl.STREAM_DRAW)

	var transform1, transform2, texcoords, texlevel gl.AttribLocation
	transform1 = 0
	transform2 = 1
	texcoords = 2
	texlevel = 3

	transform1.AttribPointer(3, gl.FLOAT, false, stride, unsafe.Offsetof(sprites[0].Transform))
	transform2.AttribPointer(3, gl.FLOAT, false, stride, unsafe.Offsetof(sprites[0].Transform)+unsafe.Sizeof(sprites[0].Transform[0])*3)
	texcoords.AttribPointer(4, gl.FLOAT, false, stride, unsafe.Offsetof(sprites[0].TextureLeft))
	texlevel.AttribPointer(1, gl.FLOAT, false, stride, unsafe.Offsetof(sprites[0].Layer))

	transform1.EnableArray()
	transform2.EnableArray()
	texcoords.EnableArray()
	texlevel.EnableArray()

	gl.DrawArrays(gl.POINTS, 0, len(sprites))

	transform1.DisableArray()
	transform2.DisableArray()
	texcoords.DisableArray()
	texlevel.DisableArray()
}
