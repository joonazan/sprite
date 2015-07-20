package sprite

import (
	"fmt"
	"github.com/go-gl-legacy/gl"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(fn string) string {
	if contents, err := ioutil.ReadFile(fn); err == nil {
		return string(contents)
	} else {
		fmt.Println(err)
		return ""
	}
}

func CreateProgram(filenames ...string) gl.Program {

	program := gl.CreateProgram()

	for _, fn := range filenames {

		var shaderType gl.GLenum

		switch {
		case strings.HasSuffix(fn, ".vert"):
			shaderType = gl.VERTEX_SHADER
		case strings.HasSuffix(fn, ".geom"):
			shaderType = gl.GEOMETRY_SHADER
		case strings.HasSuffix(fn, ".frag"):
			shaderType = gl.FRAGMENT_SHADER
		default:
			fmt.Println("Wrong suffix: " + fn)
		}

		shader := gl.CreateShader(shaderType)
		defer shader.Delete()

		shader.Source(ReadFile(fn))
		shader.Compile()
		if info := shader.GetInfoLog(); info != "" {
			log.Fatal(info)
		}

		program.AttachShader(shader)
		defer program.DetachShader(shader)
	}

	program.Link()
	if info := program.GetInfoLog(); info != "" {
		log.Fatal(info)
	}

	return program
}
