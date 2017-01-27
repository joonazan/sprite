package sprite

import (
	"github.com/go-gl-legacy/gl"
	glfw "github.com/go-gl/glfw3/v3.0/glfw"
	"log"
	"runtime"
)

func OpenWindow(width, height int, caption string) *glfw.Window {
	// OpenGL haluaa että sitä käytetään aina samasta threadista
	// muuten tulee satunnaisia segfaultteja
	runtime.LockOSThread()

	glfw.Init()

	// luodaan ikkuna
	w, err := glfw.CreateWindow(width, height, caption, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	w.MakeContextCurrent()

	gl.Init()

	// Jotta painalluksista kerrottaisiin vaikka ne olisivat jo päättyneet
	w.SetInputMode(glfw.StickyKeys, glfw.True)
	w.SetInputMode(glfw.StickyMouseButtons, glfw.True)

	// läpinäkyvyys päälle (non-premultiplied alpha)
	gl.Enable(gl.BLEND)
	gl.BlendEquationSeparate(gl.FUNC_ADD, gl.FUNC_ADD)
	gl.BlendFuncSeparate(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ZERO)

	return w
}
