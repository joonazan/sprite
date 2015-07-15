package sprite

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"image"
	"testing"
	"vec2"
)

type Testbed struct {
	window *glfw.Window
	image  *image.NRGBA
}

func NewTestbed(title string) *Testbed {
	t := new(Testbed)

	t.window = OpenWindow(500, 500, title)
	glfw.SwapInterval(1)

	t.image = LoadPNG("testimage.png").(*image.NRGBA)

	return t
}

func (t *Testbed) UntilClose(drawingFunction func()) {

	for !t.window.ShouldClose() {

		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)
		drawingFunction()
		t.window.SwapBuffers()
	}

	t.window.Destroy()
}

func TestWindow(_ *testing.T) {

	t := NewTestbed("displays test image")

	t.UntilClose(func() {
		gl.DrawPixels(t.image.Bounds().Dx(), t.image.Bounds().Dy(), gl.RGBA, gl.UNSIGNED_BYTE, t.image.Pix)
	})
}

func TestSpriteDrawer(_ *testing.T) {

	t := NewTestbed("displays rotating test image")

	spritedrawer := NewSpriteDrawer(t.window, 1)

	texcoord := Upload(t.image, spritedrawer.Texture, 0, 0, 0)
	sprite := Sprite{Image: texcoord}

	x := 0.0

	t.UntilClose(func() {

		transform := vec2.Translation(vec2.Vector{-float64(t.image.Rect.Max.X) / 2, -float64(t.image.Rect.Max.Y) / 2})
		transform = vec2.Rotation(x).Mul(transform)
		x += 0.01
		sprite.Transform = transform.To32()

		spritedrawer.Draw([]Sprite{sprite})
	})
}
